package kernelmanager_worker

import (
	kernelmanagerpb "go.resf.org/peridot/tools/kernelmanager/pb"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var w Worker

// TriggerKernelUpdateWorkflow is a workflow that triggers a kernel update.
// Depending on the kernel configuration it executes different actions.
// Currently only the Repack strategy is implemented and it stops after wards.
func TriggerKernelUpdateWorkflow(ctx workflow.Context, req *kernelmanagerpb.TriggerKernelUpdateRequest) (*kernelmanagerpb.TriggerKernelUpdateResponse, error) {
	startTime := workflow.Now(ctx)

	// Get the kernel configuration
	var kernel kernelmanagerpb.Kernel
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 25 * time.Second,
	})
	err := workflow.ExecuteActivity(ctx, w.GetKernel, req.Name).Get(ctx, &kernel)
	if err != nil {
		return nil, err
	}

	// Verify the kernel configuration
	if kernel.Config == nil {
		return nil, status.Error(codes.InvalidArgument, "kernel config is nil")
	}
	if kernel.Pkg == "" {
		return nil, status.Error(codes.InvalidArgument, "kernel package is empty")
	}
	if kernel.Config.RepackOptions == nil {
		return nil, status.Error(codes.InvalidArgument, "repack options are nil")
	}

	// Repack the kernel
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 25 * time.Minute,
	})
	var update kernelmanagerpb.Update
	err = workflow.ExecuteActivity(ctx, w.KernelRepack, &kernel).Get(ctx, &update)
	if err != nil {
		return nil, err
	}

	// Set the start time
	update.StartedTime = timestamppb.New(startTime)

	return &kernelmanagerpb.TriggerKernelUpdateResponse{
		Update: &update,
	}, nil
}
