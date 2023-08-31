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

package base

import (
	"context"
	"crypto/tls"
	"github.com/pkg/errors"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"strings"
	"time"
)

type temporalTQInterceptor struct {
	interceptor.ClientInterceptorBase
	interceptor.ClientOutboundInterceptorBase

	taskQueue string
}

func (tqi *temporalTQInterceptor) InterceptClient(next interceptor.ClientOutboundInterceptor) interceptor.ClientOutboundInterceptor {
	return &temporalTQInterceptor{
		ClientOutboundInterceptorBase: interceptor.ClientOutboundInterceptorBase{
			Next: next,
		},
		taskQueue: tqi.taskQueue,
	}
}

func (tqi *temporalTQInterceptor) ExecuteWorkflow(ctx context.Context, in *interceptor.ClientExecuteWorkflowInput) (client.WorkflowRun, error) {
	in.Options.TaskQueue = tqi.taskQueue
	return tqi.Next.ExecuteWorkflow(ctx, in)
}

func NewTemporalClient(host string, namespace string, taskQueue string, opts client.Options) (client.Client, error) {
	// If host contains :443, then use TLS
	if strings.Contains(host, ":443") {
		opts.ConnectionOptions = client.ConnectionOptions{
			DialOptions: []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
			},
		}
	}

	opts.HostPort = host

	LogInfof("Using Temporal namespace %s", namespace)

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

	// Set interceptor to set task queue
	if opts.Interceptors == nil {
		opts.Interceptors = []interceptor.ClientInterceptor{}
	}
	opts.Interceptors = append(opts.Interceptors, &temporalTQInterceptor{
		taskQueue: taskQueue,
	})

	LogInfof("Connecting to Temporal at %s", host)

	// Dial Temporal
	cl, err := client.Dial(opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial temporal")
	}

	return cl, nil
}
