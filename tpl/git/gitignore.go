// Copyright 2023 huija
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

package git

import "github.com/taouniverse/taogo/constant"

// GitIgnoreTpl of .gitignore
const GitIgnoreTpl = `
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with ` + constant.BackQuote + `go test -c` + constant.BackQuote + `
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
coverage.html
coverage.txt

# Dependency directories (remove the comment below to include it)
# vendor/

# others
*.log
`
