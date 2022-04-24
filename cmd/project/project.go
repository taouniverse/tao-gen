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
	"fmt"
	"github.com/spf13/cobra"
)

var (
	author string
	module string

	Cmd = &cobra.Command{
		Use:   "project",
		Short: "Generate project based on tao universe",
		Long:  `Generate project based on tao universe, e.g. https://github.com/taouniverse/hello`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("project info, author: '%s', module: '%s'\n", author, module)
		},
	}
)

func init() {
	// Persistence Flags
	Cmd.PersistentFlags().StringVarP(&module, "module", "m", "github.com/taouniverse/tao-hello", "target module name")
	Cmd.PersistentFlags().StringVarP(&author, "author", "a", "huija", "author of the target project")
}
