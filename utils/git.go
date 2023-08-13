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
	"os"
	"os/exec"
)

// Init git init
func Init(dir string) error {
	command := exec.Command("git", "init")
	command.Dir = dir
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}

// RemoveDir just like .git
func RemoveDir(dir string) error {
	command := exec.Command("rm", "-rf", dir)
	command.Dir = "./"
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}
