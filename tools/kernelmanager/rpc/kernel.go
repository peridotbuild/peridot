package kernelmanager_rpc

import (
	"context"
	"errors"
	"fmt"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/kv"
	kernelmanagerpb "go.resf.org/peridot/tools/kernelmanager/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"strings"
)

func (s *Server) ListKernels(ctx context.Context, req *kernelmanagerpb.ListKernelsRequest) (*kernelmanagerpb.ListKernelsResponse, error) {
	// Min page size is 1, max page size is 100.
	if req.PageSize < 1 {
		req.PageSize = 20
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	prefix := fmt.Sprintf("/kernels/entries/")
	query, err := s.kv.RangePrefix(ctx, prefix, req.PageSize, req.PageToken)
	if err != nil {
		base.LogErrorf("failed to get kernels: %v", err)
		return nil, status.Error(codes.Internal, "failed to get kernels")
	}

	var kernels []*kernelmanagerpb.Kernel
	for _, pair := range query.Pairs {
		kernel := &kernelmanagerpb.Kernel{}
		err := proto.Unmarshal(pair.Value, kernel)
		if err != nil {
			base.LogErrorf("failed to unmarshal kernel: %v", err)
			return nil, status.Error(codes.Internal, "failed to unmarshal kernel")
		}
		kernels = append(kernels, kernel)
	}

	return &kernelmanagerpb.ListKernelsResponse{
		Kernels:       kernels,
		NextPageToken: query.NextToken,
	}, nil
}

func (s *Server) GetKernel(ctx context.Context, req *kernelmanagerpb.GetKernelRequest) (*kernelmanagerpb.Kernel, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name must be provided")
	}

	kernelBytes, err := s.kv.Get(ctx, fmt.Sprintf("/kernels/entries/%s", strings.TrimPrefix(req.Name, "kernels/")))
	if err != nil {
		if errors.Is(err, kv.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "kernel not found")
		}
		base.LogErrorf("failed to get kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to get kernel")
	}

	kernel := &kernelmanagerpb.Kernel{}
	err = proto.Unmarshal(kernelBytes.Value, kernel)
	if err != nil {
		base.LogErrorf("failed to unmarshal kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to unmarshal kernel")
	}

	return kernel, nil
}

func (s *Server) CreateKernel(ctx context.Context, req *kernelmanagerpb.CreateKernelRequest) (*kernelmanagerpb.Kernel, error) {
	// Authenticate the request
	if _, err := base.UserFromContext(ctx); err != nil {
		return nil, err
	}

	if req.Kernel == nil {
		return nil, status.Error(codes.InvalidArgument, "kernel must be provided")
	}

	// Verify first that the custom name is not already taken
	prefix := fmt.Sprintf("/kernels/entries/%s/", req.Kernel.Name)
	query, err := s.kv.RangePrefix(ctx, prefix, 1, "")
	if err != nil {
		base.LogErrorf("failed to get kernels: %v", err)
		return nil, status.Error(codes.Internal, "failed to get kernels")
	}
	if len(query.Pairs) > 0 {
		return nil, status.Error(codes.InvalidArgument, "kernel name already taken")
	}

	name := fmt.Sprintf("%s/%s", req.Kernel.Name, base.NameGen("kernels"))
	req.Kernel.Name = name

	kernelBytes, err := proto.Marshal(req.Kernel)
	if err != nil {
		base.LogErrorf("failed to marshal kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to marshal kernel")
	}

	err = s.kv.Set(ctx, fmt.Sprintf("/kernels/entries/%s", name), kernelBytes)
	if err != nil {
		base.LogErrorf("failed to set kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to set kernel")
	}

	return req.Kernel, nil
}

func (s *Server) UpdateKernel(ctx context.Context, req *kernelmanagerpb.UpdateKernelRequest) (*kernelmanagerpb.Kernel, error) {
	// Authenticate the request
	if _, err := base.UserFromContext(ctx); err != nil {
		return nil, err
	}

	if req.Kernel == nil {
		return nil, status.Error(codes.InvalidArgument, "kernel must be provided")
	}

	// Check existing kernel
	_, err := s.kv.Get(ctx, fmt.Sprintf("/kernels/entries/%s", req.Kernel.Name))
	if err != nil {
		if errors.Is(err, kv.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "kernel not found")
		}
		base.LogErrorf("failed to get kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to get kernel")
	}

	kernelBytes, err := proto.Marshal(req.Kernel)
	if err != nil {
		base.LogErrorf("failed to marshal kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to marshal kernel")
	}

	err = s.kv.Set(ctx, fmt.Sprintf("/kernels/entries/%s", req.Kernel.Name), kernelBytes)
	if err != nil {
		base.LogErrorf("failed to set kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to set kernel")
	}

	return req.Kernel, nil
}
