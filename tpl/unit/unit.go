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

package unit

// Unit ${unit}.go
const Unit = `
package {{ .Module | packageName }}

import (
	"github.com/taouniverse/tao"
	// Load the required dependencies.
	// An error occurs when there was no package in the root directory.
	{{ .Require | import }}
)

/**
import _ "{{ .Module }}"
*/

// {{ .Module | packageName | firstChar | toUpper }} config of {{ .Module | packageName }}
var {{ .Module | packageName | firstChar | toUpper }} = new(Config)

func init() {
	err := tao.Register(ConfigKey, {{ .Module | packageName | firstChar | toUpper }}, setup)
	if err != nil {
		panic(err.Error())
	}
}

// TODO setup unit with the global config '{{ .Module | packageName | firstChar | toUpper }}'
// execute when init tao universe
func setup() error {
	return nil
}`

// UnitTest ${unit}_test.go
const UnitTest = `
package {{ .Module | packageName }}

import (
	"github.com/stretchr/testify/assert"
	"github.com/taouniverse/tao"
	"testing"
)

func TestTao(t *testing.T) {
	err := tao.DevelopMode()
	assert.Nil(t, err)

	assert.Equal(t, {{ .Module | packageName | firstChar | toUpper }}, default{{ .Module | packageName | title }})

	err = tao.Run(nil, nil)
	assert.Nil(t, err)
}
`
