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
	parsedURI, err := url.Parse(ctx.String(string(base.EnvVarStorageConnectionString)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse storage connection string")
	}

	// Retrieve the bucket name
	bucket := parsedURI.Path

	// Remove the leading/trailing slashes
	bucket = strings.TrimSuffix(strings.TrimPrefix(bucket, "/"), "/")

	// Convert certain flags into environment variables so that they can be used by the AWS SDK
	base.RareUseChangeDefault("AWS_REGION", ctx.String(string(base.EnvVarStorageRegion)))
	base.RareUseChangeDefault("AWS_ENDPOINT", ctx.String(string(base.EnvVarStorageEndpoint)))

	if !ctx.Bool(string(base.EnvVarStorageSecure)) {
		base.RareUseChangeDefault("AWS_DISABLE_SSL", "true")
	}

	if ctx.Bool(string(base.EnvVarStoragePathStyle)) {
		base.RareUseChangeDefault("AWS_S3_FORCE_PATH_STYLE", "true")
	}

	return New(bucket)
}
