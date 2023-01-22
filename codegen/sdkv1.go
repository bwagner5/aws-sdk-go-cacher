/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/samber/lo"
)

var header = `/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/`

type Method struct {
	Name string
	// Inputs is a mapping of input name to type
	Inputs []Arg
	// Outputs is a slice of types
	Outputs []string
}

func (m Method) FindInputType() (string, bool) {
	for _, input := range m.Inputs {
		if strings.HasSuffix(input.Type, "Input") {
			return input.Type, true
		}
	}
	return "", false
}

func (m Method) FindOutputType() (string, bool) {
	for _, output := range m.Outputs {
		if strings.HasSuffix(output, "Output") {
			return output, true
		}
	}
	return "", false
}

func (m Method) CallString() string {
	inputs := lo.Map(m.Inputs, func(a Arg, _ int) string {
		if strings.HasPrefix(a.Type, "...") {
			return fmt.Sprintf("%s...", a.Name)
		}
		return a.Name
	})
	return fmt.Sprintf("%s(%s)", m.Name, strings.Join(inputs, ", "))
}

func (m Method) String() string {
	inputs := lo.Map(m.Inputs, func(a Arg, _ int) string { return a.String() })
	return fmt.Sprintf("%s(%s) (%s)", m.Name, strings.Join(inputs, ", "), strings.Join(m.Outputs, ", "))
}

type Arg struct {
	Name string
	Type string
}

func (a Arg) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

func main() {
	src := &bytes.Buffer{}
	fmt.Fprintln(src, header)
	fmt.Fprintln(src, "package imds")
	fmt.Fprintln(src, "// DO NOT EDIT")
	fmt.Fprintln(src, "// THIS FILE IS AUTO GENERATED")
	fmt.Fprintln(src, `import (
		"context"
		"fmt"
		"strconv"
		)`)

	t := reflect.TypeOf((*ec2iface.EC2API)(nil)).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		method := DescribeMethod(t.Method(i))
		if _, ok := method.FindOutputType(); ok && !strings.HasSuffix(method.Name, "Request") {
			fmt.Fprintf(src, "func (c *Client) %s {\n", method.String())
			fmt.Fprintln(src, cachableFuncBody(method))
			fmt.Fprintln(src, "}")
		}

		// TODO: write function body that
		// 1. Hashes the input
		// 2. Looks up from cache
		// 3. Passes through to client if a cache miss occurs
		// 4. Provide a registration mechanism to keep an input query up-to-date by running go routines
	}
	formatted, err := format.Source(src.Bytes())
	if err != nil {
		log.Fatalf("formatting generated source, %s", err)
	}

	fmt.Println(string(formatted))
}

func cachableFuncBody(m Method) string {
	src := &bytes.Buffer{}
	fmt.Fprintln(src, `hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})`)
	fmt.Fprintln(src, "if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {")
	fmt.Fprintf(src, "   return cachedOutput.(%s), nil\n", m.Outputs[0])
	fmt.Fprintln(src, "}")
	fmt.Fprintf(src, "out, err := c.client.%s\n", m.CallString())
	fmt.Fprintln(src, "if err != nil { return nil, err }")
	fmt.Fprintln(src, "cache.SetWithDefault(out)")
	fmt.Fprintln(src, "return out, err")
	return src.String()
}

func DescribeMethod(method reflect.Method) Method {
	m := Method{Name: method.Name}
	withCtx := false
	pages := false
	if strings.HasSuffix(method.Name, "WithContext") {
		withCtx = true
	}
	if strings.HasSuffix(method.Name, "Pages") {
		pages = true
	}

	for j := 0; j < method.Type.NumIn(); j++ {
		if withCtx && j == 0 {
			m.Inputs = append(m.Inputs, Arg{Name: "ctx", Type: method.Type.In(j).String()})
		} else if pages && j == method.Type.NumIn()-2 {
			m.Inputs = append(m.Inputs, Arg{Name: "fn", Type: method.Type.In(j).String()})
		} else if method.Type.IsVariadic() && j == method.Type.NumIn()-1 {
			m.Inputs = append(m.Inputs, Arg{Name: "opts", Type: strings.Replace(method.Type.In(j).String(), "[]", "...", 1)})
		} else {
			m.Inputs = append(m.Inputs, Arg{Name: "input", Type: method.Type.In(j).String()})
		}
	}

	for j := 0; j < method.Type.NumOut(); j++ {
		m.Outputs = append(m.Outputs, method.Type.Out(j).String())
	}
	return m
}
