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
	"archive/tar"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

var (
	interfaceRE = regexp.MustCompile(`type ([0-9a-zA-Z]+) interface`)
)

func main() {
	outputDir := flag.String("out-dir", "codegen/v1", "directory to output generated caching clients")
	flag.Parse()

	src := &bytes.Buffer{}
	fmt.Fprintln(src, header)
	fmt.Fprintln(src, "package main")
	fmt.Fprintln(src, "// DO NOT EDIT")
	fmt.Fprintln(src, "// THIS FILE IS AUTO GENERATED")
	fmt.Fprintln(src, `import (
		"log"
		"flag"
        "fmt"
        "os"
        "errors"
        `)
	services := getServices()
	for _, service := range services {
		serviceShort := strings.ToLower(ReplaceLast(strings.Split(service, ".")[1], "API", ""))
		serviceIfaceShort := strings.ToLower(strings.Split(service, ".")[0])
		fmt.Fprintf(src, `"github.com/aws/aws-sdk-go/service/%s/%s"
        `, serviceShort, serviceIfaceShort)
	}
	fmt.Fprintln(src, ")")
	fmt.Fprintln(src, "func main() {")
	fmt.Fprintln(src, `
    outputDir := flag.String("out-dir", "pkg", "directory to output generated caching clients")
    flag.Parse()
    `)
	fmt.Fprintln(src, `var serviceOutDir string
    var err error
    var out string`)

	for _, service := range services {
		serviceShort := strings.ToLower(ReplaceLast(strings.Split(service, ".")[1], "API", ""))
		fmt.Fprintf(src, `out, err = GenSDK[%s]()
        if err != nil {
            log.Printf("%%s: %%v\n", "%s.go", err)
        }
        serviceOutDir = fmt.Sprintf("%%s/%scacher", *outputDir)
        if out != "" {
            if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
                err := os.Mkdir(serviceOutDir, os.ModePerm)
                if err != nil {
                    log.Fatal(err)
                }
            }
            os.WriteFile(fmt.Sprintf("%%s/%s.go", serviceOutDir), []byte(out), 0644)
        }
        `,
			service, fmt.Sprintf("%sapi", serviceShort), serviceShort, fmt.Sprintf("%scacher", serviceShort))
	}
	fmt.Fprintln(src, "}")
	formatted, err := format.Source(src.Bytes())
	if err != nil {
		fmt.Println(src.String())
		log.Fatalf("formatting generated source, %s", err)
	}
	os.WriteFile(fmt.Sprintf("%s/main.go", *outputDir), formatted, 0644)
}

func getServices() []string {
	resp, err := http.Get("https://github.com/aws/aws-sdk-go/archive/refs/tags/v1.44.189.tar.gz")
	if err != nil {
		log.Fatalf("Unable to download the aws-sdk-go v1: %s", err)
	}
	defer resp.Body.Close()
	temp := os.TempDir()
	if err := Untar(temp, resp.Body); err != nil {
		log.Fatalf("Unable to untar aws-sdk-go v1: %s", err)
	}
	matches, err := filepath.Glob(fmt.Sprintf("%s/aws-sdk-go-*/service", temp))
	if err != nil {
		log.Fatal(err)
	}
	if len(matches) == 0 {
		log.Fatalf("no matches found for aws-sdk-go-* in %s", temp)
	}
	if len(matches) > 1 {
		log.Fatalf("too many matches found (%d) for aws-sdk-go-* in %s", len(matches), temp)
	}
	servicesDir := matches[0]
	files, err := os.ReadDir(servicesDir)
	if err != nil {
		log.Fatal(err)
	}
	var services []string
	for _, file := range files {
		if file.IsDir() {
			serviceFiles, err := os.ReadDir(fmt.Sprintf("%s/%s", servicesDir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			for _, serviceFile := range serviceFiles {
				if serviceFile.IsDir() && strings.Contains(serviceFile.Name(), "iface") {
					ifaceFile, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/interface.go", servicesDir, file.Name(), serviceFile.Name()))
					if err != nil {
						log.Fatal(err)
					}
					submatches := interfaceRE.FindStringSubmatch(string(ifaceFile))
					if len(submatches) == 0 {
						log.Fatal("can't find interface name")
					}
					services = append(services, fmt.Sprintf("%s.%s", serviceFile.Name(), submatches[1]))
				}
			}
		}
	}
	return services
}

func ReplaceLast(str string, old string, replacement string) string {
	return reverse(strings.Replace(reverse(str), reverse(old), reverse(replacement), 1))
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
