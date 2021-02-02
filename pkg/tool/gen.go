// Copyright 2019 CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

// TODO: remove when we have a cuedoc server. Until then,
// piggyback on pkg.go.dev.

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const msg = `// Code generated by cue get go. DO NOT EDIT.

// Package tool defines statefull operation types for cue commands.
//
// This package is only visible in cue files with a _tool.cue or _tool_test.cue
// ending.
//
// CUE configuration files are not influenced by and do not influence anything
// outside the configuration itself: they are hermetic. Tools solve
// two problems: allow outside values such as environment variables,
// file or web contents, random generators etc. to influence configuration,
// and allow configuration to be actionable from within the tooling itself.
// Separating these concerns makes it clear to user when outside influences are
// in play and the tool definition can be strict about what is allowed.
//
// Tools are defined in files ending with _tool.cue. These files have a
// top-level map, "command", which defines all the tools made available through
// the cue command.
//
// The following definitions are for defining commands in tool files:
//     %s
package tool
`

func main() {
	f, _ := os.Create("doc.go")
	defer f.Close()
	b, _ := ioutil.ReadFile("tool.cue")
	i := bytes.Index(b, []byte("package tool"))
	b = b[i+len("package tool")+1:]
	b = bytes.ReplaceAll(b, []byte("\n"), []byte("\n//     "))
	fmt.Fprintf(f, msg, string(b))
}
