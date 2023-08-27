package base

import (
	"context"
	"crypto/tls"
	"github.com/pkg/errors"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"strings"
	"time"
)

func NewTemporalClient(host string, namespace string, opts client.Options) (client.Client, error) {
	// If host contains :443, then use TLS
	if strings.Contains(host, ":443") {
		opts.ConnectionOptions = client.ConnectionOptions{
			DialOptions: []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
			},
		}
	}

	LogInfof("Connecting to Temporal at %s", host)

	opts.HostPort = host

	// Register namespace (ignore error if already exists)
	nscl, err := client.NewNamespaceClient(opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create namespace client")
	}

	// Set default retention period to 5 days
	dur := 5 * 24 * time.Hour
	err = nscl.Register(context.TODO(), &workflowservice.RegisterNamespaceRequest{
		Namespace:                        namespace,
		WorkflowExecutionRetentionPeriod: &dur,
	})
	if err != nil && !strings.Contains(err.Error(), "Namespace already exists") {
		return nil, errors.Wrap(err, "failed to register namespace")
	}

	// Set namespace in opts
	opts.Namespace = namespace

	// Dial Temporal
	cl, err := client.Dial(opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial temporal")
	}

	return cl, nil
}
