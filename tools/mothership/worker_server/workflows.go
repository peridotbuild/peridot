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

// processRPMPostHold is a part of the ProcessRPM workflow.
// This part executes the import part, and retries if it fails.
// After the first failure, the workflow is put on hold.
// If the workflow is put on hold, the workflow can be rescued by an admin.
func processRPMPostHold(ctx workflow.Context, entry *mothershippb.Entry, args *mothershippb.ProcessRPMArgs) (*mothershippb.ProcessRPMResponse, error) {
	// If resource exists, then we can start the import.
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		// We'll wait up to 5 minutes for the import to finish.
		// Most imports are fast, but some packages are very large.
		StartToCloseTimeout: 5 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})
	var importRpmRes mothershippb.ImportRPMResponse
	err := workflow.ExecuteActivity(ctx, w.ImportRPM, args.Request.RpmUri, args.Request.Checksum, args.Request.OsRelease).Get(ctx, &importRpmRes)
	if err != nil {
		// If the import fails, we'll put the workflow on hold.
		// If the workflow is put on hold, an admin can rescue the workflow.
		var err error
		signalChan := workflow.GetSignalChannel(ctx, "rescue")
		workflow.GetLogger(ctx).Info("Import failed, putting workflow on hold")
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(ctx.Done(), func(c workflow.ReceiveChannel, more bool) {
			err = ctx.Err()
		})
		selector.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, nil)
		})

		// Set state to on hold
		ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 25 * time.Second,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 0,
			},
		})
		err = workflow.ExecuteActivity(ctx, w.SetEntryState, entry.Name, mothershippb.Entry_ON_HOLD, nil).Get(ctx, nil)
		if err != nil {
			return nil, err
		}

		// Wait until a rescue signal is received. Otherwise, an admin can also
		// cancel the workflow.
		selector.Select(ctx)

		// Check if workflow was cancelled.
		if err != nil {
			ctx, cancel := workflow.NewDisconnectedContext(ctx)
			defer cancel()
			ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
				StartToCloseTimeout: 25 * time.Second,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 0,
				},
			})
			_ = workflow.ExecuteActivity(ctx, w.SetEntryState, entry.Name, mothershippb.Entry_CANCELLED, nil).Get(ctx, nil)
			return nil, err
		}

		// Set the entry state to archiving
		ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 25 * time.Second,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 0,
			},
		})
		err = workflow.ExecuteActivity(ctx, w.SetEntryState, entry.Name, mothershippb.Entry_ARCHIVING, nil).Get(ctx, nil)
		if err != nil {
			return nil, err
		}

		// If the workflow was not cancelled, then we can retry the import.
		return processRPMPostHold(ctx, entry, args)
	}

	// If the import succeeds, then we can update the entry state.
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 25 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 0,
		},
	})
	err = workflow.ExecuteActivity(ctx, w.SetEntryState, entry.Name, mothershippb.Entry_ARCHIVED, &importRpmRes).Get(ctx, entry)
	if err != nil {
		return nil, err
	}

	return &mothershippb.ProcessRPMResponse{
		Entry: entry,
	}, nil
}

// ProcessRPMWorkflow processes an SRPM.
// Usually a client worker will first initiate an upload to the storage backend,
// then send a request to the Server `SubmitEntry` method (or send a request
// then upload the resource).
func ProcessRPMWorkflow(ctx workflow.Context, args *mothershippb.ProcessRPMArgs) (*mothershippb.ProcessRPMResponse, error) {
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
	err := workflow.ExecuteActivity(ctx, w.VerifyResourceExists, args.Request.RpmUri).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Set worker last check in time
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 25 * time.Second,
	})
	err = workflow.ExecuteActivity(ctx, w.SetWorkerLastCheckinTime, args.InternalRequest.WorkerId).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Create an entry, if the import fails, we'll still have an entry.
	// If it succeeds, we'll update the entry state.
	// If it fails we can set the workflow on hold and if the patches are updated
	// an admin can signal and "rescue" the workflow.
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 25 * time.Second,
	})
	var entry mothershippb.Entry
	err = workflow.ExecuteActivity(ctx, w.CreateEntry, args).Get(ctx, &entry)
	if err != nil {
		return nil, err
	}

	// On defer, if the workflow is not completed, then we'll set the entry state
	// to failed.
	defer func() {
		if entry.State == mothershippb.Entry_ARCHIVED {
			return
		}

		ctx, cancel := workflow.NewDisconnectedContext(ctx)
		defer cancel()
		ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 25 * time.Second,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 0,
			},
		})
		_ = workflow.ExecuteActivity(ctx, w.SetEntryState, entry.Name, mothershippb.Entry_FAILED, nil).Get(ctx, nil)
	}()

	// Set the entry name to the RPM NVR
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 45 * time.Second,
	})
	err = workflow.ExecuteActivity(ctx, w.SetEntryIDFromRPM, entry.Name, args.Request.RpmUri, args.Request.Checksum).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Process the RPM.
	return processRPMPostHold(ctx, &entry, args)
}
