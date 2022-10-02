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
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTemplate(t *testing.T) {
	assert.Equal(t, importFunc(""), "")
	assert.Equal(t, modRequire(""), "")

	r, err := packageName("github.com/tao/tao-hello")
	assert.Nil(t, err)
	assert.Equal(t, r, "hello")

	r, err = firstChar(r)
	assert.Nil(t, err)
	assert.Equal(t, r, "h")

	assert.Equal(t, strings.Title("hello"), "Hello")

	r, err = modVersion("go1.17.2")
	assert.Nil(t, err)
	assert.Equal(t, r, "go 1.17")
}
