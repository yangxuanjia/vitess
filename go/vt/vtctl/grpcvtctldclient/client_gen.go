// Code generated by grpcvtctldclient-generator. DO NOT EDIT.

/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpcvtctldclient

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	vtctldatapb "vitess.io/vitess/go/vt/proto/vtctldata"
)

// FindAllShardsInKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) FindAllShardsInKeyspace(ctx context.Context, in *vtctldatapb.FindAllShardsInKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.FindAllShardsInKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.FindAllShardsInKeyspace(ctx, in, opts...)
}

// GetBackups is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetBackups(ctx context.Context, in *vtctldatapb.GetBackupsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetBackupsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetBackups(ctx, in, opts...)
}

// GetCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetCellInfo(ctx context.Context, in *vtctldatapb.GetCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetCellInfo(ctx, in, opts...)
}

// GetCellInfoNames is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetCellInfoNames(ctx context.Context, in *vtctldatapb.GetCellInfoNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoNamesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetCellInfoNames(ctx, in, opts...)
}

// GetCellsAliases is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetCellsAliases(ctx context.Context, in *vtctldatapb.GetCellsAliasesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellsAliasesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetCellsAliases(ctx, in, opts...)
}

// GetKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetKeyspace(ctx context.Context, in *vtctldatapb.GetKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetKeyspace(ctx, in, opts...)
}

// GetKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetKeyspaces(ctx context.Context, in *vtctldatapb.GetKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspacesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetKeyspaces(ctx, in, opts...)
}

// GetSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSchema(ctx context.Context, in *vtctldatapb.GetSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSchema(ctx, in, opts...)
}

// GetSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSrvVSchema(ctx context.Context, in *vtctldatapb.GetSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSrvVSchema(ctx, in, opts...)
}

// GetTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetTablet(ctx context.Context, in *vtctldatapb.GetTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetTablet(ctx, in, opts...)
}

// GetTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetTablets(ctx context.Context, in *vtctldatapb.GetTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetTablets(ctx, in, opts...)
}

// GetVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetVSchema(ctx context.Context, in *vtctldatapb.GetVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetVSchema(ctx, in, opts...)
}

// InitShardPrimary is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) InitShardPrimary(ctx context.Context, in *vtctldatapb.InitShardPrimaryRequest, opts ...grpc.CallOption) (*vtctldatapb.InitShardPrimaryResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.InitShardPrimary(ctx, in, opts...)
}
