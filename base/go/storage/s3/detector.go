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

package storage_s3

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"net/url"
	"strings"
)

func FromFlags(ctx *cli.Context) (*S3, error) {
	// Parse the connection string
	parsedURI, err := url.Parse(ctx.String("storage-connection-string"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse storage connection string")
	}

	// Retrieve the bucket name
	bucket := parsedURI.Host

	// Remove the leading/trailing slashes
	bucket = strings.TrimSuffix(strings.TrimPrefix(bucket, "/"), "/")

	// Convert certain flags into environment variables so that they can be used by the AWS SDK
	base.RareUseChangeDefault("AWS_REGION", ctx.String("storage-region"))
	base.RareUseChangeDefault("AWS_ENDPOINT", ctx.String("storage-endpoint"))

	if !ctx.Bool("storage-secure") {
		base.RareUseChangeDefault("AWS_DISABLE_SSL", "true")
	}

	if ctx.Bool("storage-path-style") {
		base.RareUseChangeDefault("AWS_S3_FORCE_PATH_STYLE", "true")
	}

	base.LogInfof("Using S3 bucket: %s", bucket)

	return New(bucket)
}
