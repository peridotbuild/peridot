package storage_detector

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/storage"
	storage_s3 "go.resf.org/peridot/base/go/storage/s3"
	"net/url"
)

func FromFlags(ctx *cli.Context) (storage.Storage, error) {
	parsedURI, err := url.Parse(ctx.String(string(base.EnvVarStorageConnectionString)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse storage connection string")
	}

	switch parsedURI.Scheme {
	case "s3":
		return storage_s3.FromFlags(ctx)
	default:
		return nil, errors.Errorf("unknown storage scheme: %s", parsedURI.Scheme)
	}
}
