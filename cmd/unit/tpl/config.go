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

package tpl

import (
	"github.com/taouniverse/gen/cmd/utils"
)

// Config config.go
const Config = `
package {{ .Module | package }}

import (
	"context"
	"github.com/taouniverse/tao"
	{{ .Require | import }}
)

// ConfigKey for this repo
const ConfigKey = "{{ .Module | package }}"

// Config implements tao.Config
type Config struct {
	RunAfters []string ` + utils.BackQuote + `json:"run_after,omitempty"` + utils.BackQuote + `
}

var default{{ .Module | package | title }} = &Config{
	RunAfters: []string{},
}

// Default config
func (h *Config) Default() tao.Config {
	return default{{ .Module | package | title }}
}

// ValidSelf with some default values
func (h *Config) ValidSelf() {
	if h.RunAfters == nil {
		h.RunAfters = default{{ .Module | package | title }}.RunAfters
	}
}

// ToTask transform itself to Task
func (h *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			// non-block check
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			// TODO JOB
			return param, nil
		})
}

// RunAfter defines pre task names
func (h *Config) RunAfter() []string {
	return h.RunAfters
}
`

// ConfigTest config_test.go
const ConfigTest = `
package {{ .Module | package }}

import (
	"github.com/stretchr/testify/assert"
	"github.com/taouniverse/tao"
	"testing"
)

func TestTao(t *testing.T) {
	tao.DevelopMode()

	assert.Equal(t, {{ .Module | package | first | upper }}, default{{ .Module | package | title }})

	err := tao.Run(nil, nil)
	assert.Nil(t, err)
}
`
