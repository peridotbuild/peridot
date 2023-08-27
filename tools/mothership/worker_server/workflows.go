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

package mothership_worker_server

import (
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

var w Worker

// ProcessRPMWorkflow processes an SRPM.
// Usually a client worker will first initiate an upload to the storage backend,
// then send a request to the Server `SubmitEntry` method (or send a request
// then upload the resource).
func ProcessRPMWorkflow(ctx workflow.Context, req *mothershippb.ProcessRPMRequest) (*mothershippb.ProcessRPMResponse, error) {
	// First verify that the resource exists.
	// The resource can be uploaded after the request is sent.
	// So we should wait up to 2 hours. The initial timeouts should be low
	// since the worker is most likely to upload the resource immediately.
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 25 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			// We're waiting 25 seconds each time
			InitialInterval:    25 * time.Second,
			BackoffCoefficient: 1,
			// Maximum attempts should be set, so it's approximately 2 hours
			MaximumAttempts: (60 * 60 * 2) / 25,
		},
	})
	err := workflow.ExecuteActivity(ctx, w.VerifyResourceExists, req.RpmUri).Get(ctx, nil)
	if err != nil {
		return nil, err
	}
}
