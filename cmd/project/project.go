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

package project

import (
	"github.com/spf13/cobra"
	"github.com/taouniverse/tao"
	"github.com/taouniverse/taogo/tpl/license"
	"github.com/taouniverse/taogo/tpl/project"
	"github.com/taouniverse/taogo/utils"
	"strconv"
	"time"
)

var (
	author  string
	name    string
	module  string
	require string
	dir     string

	Cmd = &cobra.Command{
		Use:   "project",
		Short: "Generate project based on tao universe",
		Long:  `Generate project based on tao universe, e.g. https://github.com/taouniverse/hello`,
		Run: func(cmd *cobra.Command, args []string) {
			path, err := utils.CheckDir(dir, name)
			if err != nil {
				tao.Panic(err)
			}
			templates := map[string]string{
				path + "main.go": license.Apache2FileHeaderTpl + project.Main,
				path + "go.mod":  project.Mod,
				path + "LICENSE": license.Apache2LicenseFileTpl,
			}
			params := map[string]string{
				"Author":  author,
				"Module":  module,
				"Require": require,
				"Year":    strconv.Itoa(time.Now().Year()),
			}
			err = utils.ExecuteTemplate(templates, params)
			if err != nil {
				tao.Panic(err)
			}
			output, err := utils.ModTidy(path)
			tao.Debugf("%s", string(output))
			if err != nil {
				tao.Panic(err)
			}
		},
	}
)

func init() {
	// Persistence Flags
	Cmd.PersistentFlags().StringVarP(&module, "module", "m", "github.com/taouniverse/hello", "target module name of project")
	Cmd.PersistentFlags().StringVarP(&require, "require", "r", "github.com/taouniverse/tao-hello", "require modules, split by "+utils.Split)
	Cmd.PersistentFlags().StringVarP(&dir, "dir", "d", "./", "project's parent path")
	Cmd.PersistentFlags().StringVarP(&name, "name", "n", "hello", "name of the target project")
	Cmd.PersistentFlags().StringVarP(&author, "author", "a", "huija", "author of the target project")
}
