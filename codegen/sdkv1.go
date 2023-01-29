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
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
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

type SDK struct {
	// Name is the upper-case SDK name (i.e. EC2API)
	Name string
	// ShortName is the SDK name without the API suffix (i.e. ec2)
	ShortName string
	// ShortNameLower is the lower-cased SDK name without the API suffix (i.e. ec2)
	ShortNameLower string
}

type Method struct {
	// Name is the method name
	Name string
	// Inputs is a mapping of input name to type
	Inputs []Arg
	// Outputs is a slice of types
	Outputs []Arg
	// Pager indicates the method takes a user supplied paging function
	Pager bool
	// Context indicates the method takes a context as the first parameter
	Context bool
	// Readonly indicated the method is only reads and doesn't mutate state
	ReadOnly bool
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
		if strings.HasSuffix(output.Type, "Output") {
			return output.Type, true
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
	outputs := lo.Map(m.Outputs, func(a Arg, _ int) string { return a.Type })
	return fmt.Sprintf("%s(%s) (%s)", m.Name, strings.Join(inputs, ", "), strings.Join(outputs, ", "))
}

type Arg struct {
	Name       string
	Type       string
	Kind       reflect.Kind
	ActualType reflect.Type
}

func (a Arg) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

func main() {
	outputDir := flag.String("out-dir", "pkg/cacher", "directory to output generated caching clients")
	out, err := GenSDK[ec2iface.EC2API]()
	if err != nil {
		log.Printf("%s: %v\n", "ec2api.go", err)
	}
	os.WriteFile(fmt.Sprintf("%s/ec2api.go", *outputDir), []byte(out), 0644)

	out, err = GenSDK[s3iface.S3API]()
	if err != nil {
		log.Printf("%s: %v\n", "s3api.go", err)
	}
	os.WriteFile(fmt.Sprintf("%s/s3api.go", *outputDir), []byte(out), 0644)

	out, err = GenSDK[ssmiface.SSMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmapi.go", err)
	}
	os.WriteFile(fmt.Sprintf("%s/ssmapi.go", *outputDir), []byte(out), 0644)

	out, err = GenSDK[autoscalingiface.AutoScalingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "autoscalingapi.go", err)
	}
	os.WriteFile(fmt.Sprintf("%s/autoscalingapi.go", *outputDir), []byte(out), 0644)
}

func GenSDK[T any]() (string, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	sdk := SDK{
		Name:           t.Name(),
		ShortName:      strings.ReplaceAll(t.Name(), "API", ""),
		ShortNameLower: strings.ToLower(strings.ReplaceAll(t.Name(), "API", "")),
	}
	src := &bytes.Buffer{}
	fmt.Fprintln(src, header)
	fmt.Fprintln(src, "package cacher")
	fmt.Fprintln(src, "// DO NOT EDIT")
	fmt.Fprintln(src, "// THIS FILE IS AUTO GENERATED")
	fmt.Fprintf(src, `import (
		"context"
		"strconv"
		"time"

		"github.com/aws/aws-sdk-go/service/%s/%siface"
		"github.com/aws/aws-sdk-go/service/%s"
		"github.com/aws/aws-sdk-go/aws/request"
		"github.com/imdario/mergo"
		"github.com/mitchellh/hashstructure/v2"
		"github.com/patrickmn/go-cache"
		)
		`, sdk.ShortNameLower, sdk.ShortNameLower, sdk.ShortNameLower)
	fmt.Fprintf(src, `type %s struct {
		%siface.%s
		cache  *cache.Cache
	}
	`, sdk.ShortName, sdk.ShortNameLower, sdk.Name)
	fmt.Fprintf(src, `func New%s(%sapi %siface.%s) *%s {
		return &%s {
			%s: %sapi,
			cache: cache.New(1*time.Minute, 2*time.Minute),
		}
	}
	`, sdk.ShortName, sdk.ShortNameLower, sdk.ShortNameLower, sdk.Name, sdk.ShortName, sdk.ShortName, sdk.Name, sdk.ShortNameLower)
	for i := 0; i < t.NumMethod(); i++ {
		method := DescribeMethod(sdk, t.Method(i))
		if !method.ReadOnly {
			continue
		} else if strings.HasSuffix(method.Name, "Request") {
			continue
		} else if _, ok := method.FindOutputType(); ok {
			fmt.Fprintf(src, "func (c *%s) %s {\n", sdk.ShortName, method.String())
			fmt.Fprintln(src, cachableFuncBody(sdk, method))
			fmt.Fprintln(src, "}")
		} else if method.Pager {
			fmt.Fprintf(src, "func (c *%s) %s {\n", sdk.ShortName, method.String())
			fmt.Fprintln(src, cachablePagerFuncBody(sdk, method))
			fmt.Fprintln(src, "}")
		}

		// TODO: write function body that
		// 1. Provide a registration mechanism to keep an input query up-to-date by running go routines
	}
	formatted, err := format.Source(src.Bytes())
	if err != nil {
		return src.String(), fmt.Errorf("formatting generated source, %w", err)
	}
	return string(formatted), nil
}

func cachableFuncBody(sdk SDK, m Method) string {
	return fmt.Sprintf(`hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(%s), nil
	}
	out, err := c.%s.%s
	if err != nil { return nil, err }
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err`, m.Outputs[0], sdk.Name, m.CallString())
}

func cachablePagerFuncBody(sdk SDK, m Method) string {
	var output string
	if m.Context {
		output = m.Inputs[2].ActualType.In(0).String()
	} else {
		output = m.Inputs[1].ActualType.In(0).String()
	}
	return fmt.Sprintf(`hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(%s), false)
		return nil
	}
	cachable := true
	output := &%s{}
	fnCacher := func(out %s, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		} 
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.%s.%s; err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil`, output, strings.Replace(output, "*", "", 1), output, sdk.Name, strings.Replace(m.CallString(), ", fn", ", fnCacher", 1))
}

func DescribeMethod(sdk SDK, method reflect.Method) Method {
	m := Method{Name: method.Name}
	if strings.HasSuffix(method.Name, "WithContext") {
		m.Context = true
	} else if strings.HasSuffix(method.Name, "Pages") {
		m.Pager = true
	}
	if strings.HasSuffix(method.Name, "PagesWithContext") {
		m.Pager = true
		m.Context = true
	}
	if strings.HasPrefix(method.Name, "Get") || strings.HasPrefix(method.Name, "List") || strings.HasPrefix(method.Name, "Describe") ||
		strings.HasPrefix(method.Name, "Search") {
		m.ReadOnly = true
	}
	if !m.ReadOnly && m.Pager {
		log.Printf("Detected a mutating operation with a pager %s.%s. This may be classified incorrectly.\n", sdk.ShortName, m.Name)
	}

	for j := 0; j < method.Type.NumIn(); j++ {
		if m.Context && j == 0 {
			m.Inputs = append(m.Inputs, Arg{Name: "ctx", Type: method.Type.In(j).String(), Kind: method.Type.In(j).Kind()})
		} else if m.Pager && ((!m.Context && j == method.Type.NumIn()-1) || (m.Context && j == method.Type.NumIn()-2)) {
			m.Inputs = append(m.Inputs, Arg{Name: "fn", Type: method.Type.In(j).String(), Kind: method.Type.In(j).Kind(), ActualType: method.Type.In(j)})
		} else if method.Type.IsVariadic() && j == method.Type.NumIn()-1 {
			m.Inputs = append(m.Inputs, Arg{Name: "opts", Type: strings.Replace(method.Type.In(j).String(), "[]", "...", 1), Kind: method.Type.In(j).Kind()})
		} else {
			m.Inputs = append(m.Inputs, Arg{Name: "input", Type: method.Type.In(j).String(), Kind: method.Type.In(j).Kind()})
		}
	}

	for j := 0; j < method.Type.NumOut(); j++ {
		m.Outputs = append(m.Outputs, Arg{Type: method.Type.Out(j).String(), Kind: method.Type.Out(j).Kind()})
	}
	return m
}
