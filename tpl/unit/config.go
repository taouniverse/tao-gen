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

import (
	"github.com/taouniverse/taogo/utils"
)

// Config config.go
const Config = `
package {{ .Module | packageName }}

import (
	"context"
	"github.com/taouniverse/tao"
)

// ConfigKey for this repo
const ConfigKey = "{{ .Module | packageName }}"

// Config implements tao.Config
// TODO declare the configuration you want & define some default values
type Config struct {
	RunAfters []string ` + utils.BackQuote + `json:"run_after,omitempty"` + utils.BackQuote + `
}

var default{{ .Module | packageName | title }} = &Config{
	RunAfters: []string{},
}

// Name of Config
func ({{ .Module | packageName | firstChar }} *Config) Name() string {
	return ConfigKey
}

// ValidSelf with some default values
func ({{ .Module | packageName | firstChar }} *Config) ValidSelf() {
	if {{ .Module | packageName | firstChar }}.RunAfters == nil {
		{{ .Module | packageName | firstChar }}.RunAfters = default{{ .Module | packageName | title }}.RunAfters
	}
}

// ToTask transform itself to Task
func ({{ .Module | packageName | firstChar }} *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			// non-block check
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			// TODO JOB code run after RunAfters, you can just do nothing here
			return param, nil
		})
}

// RunAfter defines pre task names
func ({{ .Module | packageName | firstChar }} *Config) RunAfter() []string {
	return {{ .Module | packageName | firstChar }}.RunAfters
}
`

// ConfigTest config_test.go
const ConfigTest = `
package {{ .Module | packageName }}

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	{{ .Module | packageName | firstChar }} := new(Config)
	{{ .Module | packageName | firstChar }}.ValidSelf()
	assert.EqualValues(t, {{ .Module | packageName | firstChar }}, default{{ .Module | packageName | title }})

	t.Log({{ .Module | packageName | firstChar }}.RunAfter())
	t.Log({{ .Module | packageName | firstChar }}.ToTask())
}
`
