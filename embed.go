// Copyright 2026 The OpenAgent Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build embed

// This file is only compiled when building with -tags embed.
// It embeds conf/, web/build/, and skills/ into
// the binary, and wires them up via embedsupport.Setup so that the server
// can run from a single executable without any on-disk assets.
// On-disk files always take priority over the embedded versions at runtime.

package main

import (
	"embed"
	"io/fs"

	"github.com/the-open-agent/openagent/embedsupport"
)

//go:embed conf
var _embeddedConf embed.FS

//go:embed web/build/index.html web/build/favicon.ico
//go:embed web/build/static
var _embeddedWeb embed.FS

//go:embed skills
var _embeddedSkills embed.FS

func init() {
	confFS, _ := fs.Sub(_embeddedConf, "conf")
	webFS, _ := fs.Sub(_embeddedWeb, "web/build")
	skillsFS, _ := fs.Sub(_embeddedSkills, "skills")
	embedsupport.Setup(confFS, webFS, skillsFS)
}
