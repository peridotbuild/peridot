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

package srpm_import

import (
	"github.com/rocky-linux/srpmproc/pkg/data"
	"github.com/rocky-linux/srpmproc/pkg/misc"
	"go.resf.org/peridot/base/go/storage"
	"strings"
)

type srpmprocBlobCompat struct {
	storage.Storage
}

func (s *srpmprocBlobCompat) Write(path string, content []byte) error {
	_, err := s.PutBytes(path, content)
	return err
}

func (s *srpmprocBlobCompat) Read(path string) ([]byte, error) {
	return s.Get(path)
}

type srpmprocImportModeCompat struct {
	data.ImportMode
}

func (s *srpmprocImportModeCompat) ImportName(pd *data.ProcessData, md *data.ModeData) string {
	if misc.GetTagImportRegex(pd).MatchString(md.TagBranch) {
		match := misc.GetTagImportRegex(pd).FindStringSubmatch(md.TagBranch)
		return match[3]
	}

	return strings.Replace(strings.TrimPrefix(md.TagBranch, "refs/heads/"), "%", "_", -1)
}
