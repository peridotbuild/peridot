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

package kernelmanager_ui

import (
	"embed"
	base "go.resf.org/peridot/base/go"
)

//go:embed *
var assets embed.FS

func InitFrontendInfo(info *base.FrontendInfo) *embed.FS {
	if info == nil {
		info = &base.FrontendInfo{}
	}
	info.Title = "RESF KernelManager"
	info.AllowUnauthenticated = true

	return &assets
}
