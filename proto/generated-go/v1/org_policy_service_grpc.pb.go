// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/org_policy_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrgPolicyServiceClient is the client API for OrgPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgPolicyServiceClient interface {
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UndeletePolicy(ctx context.Context, in *UndeletePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
}

type orgPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgPolicyServiceClient(cc grpc.ClientConnInterface) OrgPolicyServiceClient {
	return &orgPolicyServiceClient{cc}
}

func (c *orgPolicyServiceClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/bytebase.v1.OrgPolicyService/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgPolicyServiceClient) ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error) {
	out := new(ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/bytebase.v1.OrgPolicyService/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgPolicyServiceClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/bytebase.v1.OrgPolicyService/CreatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgPolicyServiceClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/bytebase.v1.OrgPolicyService/UpdatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgPolicyServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bytebase.v1.OrgPolicyService/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgPolicyServiceClient) UndeletePolicy(ctx context.Context, in *UndeletePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/bytebase.v1.OrgPolicyService/UndeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgPolicyServiceServer is the server API for OrgPolicyService service.
// All implementations must embed UnimplementedOrgPolicyServiceServer
// for forward compatibility
type OrgPolicyServiceServer interface {
	GetPolicy(context.Context, *GetPolicyRequest) (*Policy, error)
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*emptypb.Empty, error)
	UndeletePolicy(context.Context, *UndeletePolicyRequest) (*Policy, error)
	mustEmbedUnimplementedOrgPolicyServiceServer()
}

// UnimplementedOrgPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrgPolicyServiceServer struct {
}

func (UnimplementedOrgPolicyServiceServer) GetPolicy(context.Context, *GetPolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedOrgPolicyServiceServer) ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}
func (UnimplementedOrgPolicyServiceServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedOrgPolicyServiceServer) UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedOrgPolicyServiceServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedOrgPolicyServiceServer) UndeletePolicy(context.Context, *UndeletePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UndeletePolicy not implemented")
}
func (UnimplementedOrgPolicyServiceServer) mustEmbedUnimplementedOrgPolicyServiceServer() {}

// UnsafeOrgPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgPolicyServiceServer will
// result in compilation errors.
type UnsafeOrgPolicyServiceServer interface {
	mustEmbedUnimplementedOrgPolicyServiceServer()
}

func RegisterOrgPolicyServiceServer(s grpc.ServiceRegistrar, srv OrgPolicyServiceServer) {
	s.RegisterService(&OrgPolicyService_ServiceDesc, srv)
}

func _OrgPolicyService_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgPolicyServiceServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.OrgPolicyService/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgPolicyServiceServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgPolicyService_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgPolicyServiceServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.OrgPolicyService/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgPolicyServiceServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgPolicyService_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgPolicyServiceServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.OrgPolicyService/CreatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgPolicyServiceServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgPolicyService_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgPolicyServiceServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.OrgPolicyService/UpdatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgPolicyServiceServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgPolicyService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgPolicyServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.OrgPolicyService/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgPolicyServiceServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgPolicyService_UndeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgPolicyServiceServer).UndeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytebase.v1.OrgPolicyService/UndeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgPolicyServiceServer).UndeletePolicy(ctx, req.(*UndeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgPolicyService_ServiceDesc is the grpc.ServiceDesc for OrgPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.OrgPolicyService",
	HandlerType: (*OrgPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicy",
			Handler:    _OrgPolicyService_GetPolicy_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _OrgPolicyService_ListPolicies_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _OrgPolicyService_CreatePolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _OrgPolicyService_UpdatePolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _OrgPolicyService_DeletePolicy_Handler,
		},
		{
			MethodName: "UndeletePolicy",
			Handler:    _OrgPolicyService_UndeletePolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/org_policy_service.proto",
}
