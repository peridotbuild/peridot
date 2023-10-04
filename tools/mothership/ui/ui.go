// Copyright 2023 Peridot Authors
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

package mship_ui

import (
	"embed"
	base "go.resf.org/peridot/base/go"
)

//go:embed *
var assets embed.FS

//go:embed mship_gopher.png
var gopher []byte

//go:embed mship_gopher_beta.png
var gopherBeta []byte

//go:embed favicon.png
var favicon []byte

func InitFrontendInfo(info *base.FrontendInfo) *embed.FS {
	if info == nil {
		info = &base.FrontendInfo{}
	}
	info.Title = "Mship"
	info.NoAuth = true
	info.AdditionalContent = map[string][]byte{
		"/_ga/mship_gopher.png":      gopher,
		"/_ga/mship_gopher_beta.png": gopherBeta,
		"/_ga/favicon.png":           favicon,
	}

	return &assets
}
