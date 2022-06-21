// Copyright 2022 huija
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// CheckDir before ExecuteTemplate
func CheckDir(dir, name string) (path string, err error) {
	_, err = os.Stat(dir)
	if err != nil {
		return
	}
	path = dir + "/" + name + "/"
	_, err = os.Stat(path)
	if err != nil {
		// create the target dir
		err = os.Mkdir(path, os.ModeDir|0755)
	}
	// FIXME the target dir may not be empty
	return
}

// Split of array string
const Split = ","

// ExecuteTemplate to gen files
func ExecuteTemplate(templates, params map[string]string) error {
	for f, m := range templates {
		main, err := template.New(f).Funcs(templateFuncMap).Parse(strings.TrimPrefix(m, "\n"))
		if err != nil {
			return err
		}
		file, err := os.OpenFile(f, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		err = main.Execute(file, params)
		if err != nil {
			return err
		}
	}
	return nil
}

var templateFuncMap = template.FuncMap{
	"import":  importFunc,
	"require": requireFunc,
}

func importFunc(s string) (r string) {
	imports := strings.Split(s, Split)
	for i := 0; i < len(imports); i++ {
		r += fmt.Sprintf("_ \"%s\"", imports[i])
		if i != len(imports)-1 {
			r += "\n"
		}
	}
	return
}

func requireFunc(s string) (r string) {
	imports := strings.Split(s, Split)
	for i := 0; i < len(imports); i++ {
		r += fmt.Sprintf("%s latest", imports[i])
		if i != len(imports)-1 {
			r += "\n"
		}
	}
	return
}
