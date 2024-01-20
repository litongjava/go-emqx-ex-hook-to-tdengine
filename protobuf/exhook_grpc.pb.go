//------------------------------------------------------------------------------
// Copyright (c) 2020-2021 EMQ Technologies Co., Ltd. All Rights Reserved.
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
//------------------------------------------------------------------------------

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: protobuf/exhook.proto

package exhook

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HookProvider_OnProviderLoaded_FullMethodName      = "/emqx.exhook.v1.HookProvider/OnProviderLoaded"
	HookProvider_OnProviderUnloaded_FullMethodName    = "/emqx.exhook.v1.HookProvider/OnProviderUnloaded"
	HookProvider_OnClientConnect_FullMethodName       = "/emqx.exhook.v1.HookProvider/OnClientConnect"
	HookProvider_OnClientConnack_FullMethodName       = "/emqx.exhook.v1.HookProvider/OnClientConnack"
	HookProvider_OnClientConnected_FullMethodName     = "/emqx.exhook.v1.HookProvider/OnClientConnected"
	HookProvider_OnClientDisconnected_FullMethodName  = "/emqx.exhook.v1.HookProvider/OnClientDisconnected"
	HookProvider_OnClientAuthenticate_FullMethodName  = "/emqx.exhook.v1.HookProvider/OnClientAuthenticate"
	HookProvider_OnClientCheckAcl_FullMethodName      = "/emqx.exhook.v1.HookProvider/OnClientCheckAcl"
	HookProvider_OnClientSubscribe_FullMethodName     = "/emqx.exhook.v1.HookProvider/OnClientSubscribe"
	HookProvider_OnClientUnsubscribe_FullMethodName   = "/emqx.exhook.v1.HookProvider/OnClientUnsubscribe"
	HookProvider_OnSessionCreated_FullMethodName      = "/emqx.exhook.v1.HookProvider/OnSessionCreated"
	HookProvider_OnSessionSubscribed_FullMethodName   = "/emqx.exhook.v1.HookProvider/OnSessionSubscribed"
	HookProvider_OnSessionUnsubscribed_FullMethodName = "/emqx.exhook.v1.HookProvider/OnSessionUnsubscribed"
	HookProvider_OnSessionResumed_FullMethodName      = "/emqx.exhook.v1.HookProvider/OnSessionResumed"
	HookProvider_OnSessionDiscarded_FullMethodName    = "/emqx.exhook.v1.HookProvider/OnSessionDiscarded"
	HookProvider_OnSessionTakeovered_FullMethodName   = "/emqx.exhook.v1.HookProvider/OnSessionTakeovered"
	HookProvider_OnSessionTerminated_FullMethodName   = "/emqx.exhook.v1.HookProvider/OnSessionTerminated"
	HookProvider_OnMessagePublish_FullMethodName      = "/emqx.exhook.v1.HookProvider/OnMessagePublish"
	HookProvider_OnMessageDelivered_FullMethodName    = "/emqx.exhook.v1.HookProvider/OnMessageDelivered"
	HookProvider_OnMessageDropped_FullMethodName      = "/emqx.exhook.v1.HookProvider/OnMessageDropped"
	HookProvider_OnMessageAcked_FullMethodName        = "/emqx.exhook.v1.HookProvider/OnMessageAcked"
)

// HookProviderClient is the client API for HookProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HookProviderClient interface {
	OnProviderLoaded(ctx context.Context, in *ProviderLoadedRequest, opts ...grpc.CallOption) (*LoadedResponse, error)
	OnProviderUnloaded(ctx context.Context, in *ProviderUnloadedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnClientConnect(ctx context.Context, in *ClientConnectRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnClientConnack(ctx context.Context, in *ClientConnackRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnClientConnected(ctx context.Context, in *ClientConnectedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnClientDisconnected(ctx context.Context, in *ClientDisconnectedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnClientAuthenticate(ctx context.Context, in *ClientAuthenticateRequest, opts ...grpc.CallOption) (*ValuedResponse, error)
	OnClientCheckAcl(ctx context.Context, in *ClientCheckAclRequest, opts ...grpc.CallOption) (*ValuedResponse, error)
	OnClientSubscribe(ctx context.Context, in *ClientSubscribeRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnClientUnsubscribe(ctx context.Context, in *ClientUnsubscribeRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionCreated(ctx context.Context, in *SessionCreatedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionSubscribed(ctx context.Context, in *SessionSubscribedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionUnsubscribed(ctx context.Context, in *SessionUnsubscribedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionResumed(ctx context.Context, in *SessionResumedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionDiscarded(ctx context.Context, in *SessionDiscardedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionTakeovered(ctx context.Context, in *SessionTakeoveredRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnSessionTerminated(ctx context.Context, in *SessionTerminatedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnMessagePublish(ctx context.Context, in *MessagePublishRequest, opts ...grpc.CallOption) (*ValuedResponse, error)
	OnMessageDelivered(ctx context.Context, in *MessageDeliveredRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnMessageDropped(ctx context.Context, in *MessageDroppedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
	OnMessageAcked(ctx context.Context, in *MessageAckedRequest, opts ...grpc.CallOption) (*EmptySuccess, error)
}

type hookProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewHookProviderClient(cc grpc.ClientConnInterface) HookProviderClient {
	return &hookProviderClient{cc}
}

func (c *hookProviderClient) OnProviderLoaded(ctx context.Context, in *ProviderLoadedRequest, opts ...grpc.CallOption) (*LoadedResponse, error) {
	out := new(LoadedResponse)
	err := c.cc.Invoke(ctx, HookProvider_OnProviderLoaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnProviderUnloaded(ctx context.Context, in *ProviderUnloadedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnProviderUnloaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientConnect(ctx context.Context, in *ClientConnectRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnClientConnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientConnack(ctx context.Context, in *ClientConnackRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnClientConnack_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientConnected(ctx context.Context, in *ClientConnectedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnClientConnected_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientDisconnected(ctx context.Context, in *ClientDisconnectedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnClientDisconnected_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientAuthenticate(ctx context.Context, in *ClientAuthenticateRequest, opts ...grpc.CallOption) (*ValuedResponse, error) {
	out := new(ValuedResponse)
	err := c.cc.Invoke(ctx, HookProvider_OnClientAuthenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientCheckAcl(ctx context.Context, in *ClientCheckAclRequest, opts ...grpc.CallOption) (*ValuedResponse, error) {
	out := new(ValuedResponse)
	err := c.cc.Invoke(ctx, HookProvider_OnClientCheckAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientSubscribe(ctx context.Context, in *ClientSubscribeRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnClientSubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnClientUnsubscribe(ctx context.Context, in *ClientUnsubscribeRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnClientUnsubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionCreated(ctx context.Context, in *SessionCreatedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionCreated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionSubscribed(ctx context.Context, in *SessionSubscribedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionSubscribed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionUnsubscribed(ctx context.Context, in *SessionUnsubscribedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionUnsubscribed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionResumed(ctx context.Context, in *SessionResumedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionResumed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionDiscarded(ctx context.Context, in *SessionDiscardedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionDiscarded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionTakeovered(ctx context.Context, in *SessionTakeoveredRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionTakeovered_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnSessionTerminated(ctx context.Context, in *SessionTerminatedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnSessionTerminated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnMessagePublish(ctx context.Context, in *MessagePublishRequest, opts ...grpc.CallOption) (*ValuedResponse, error) {
	out := new(ValuedResponse)
	err := c.cc.Invoke(ctx, HookProvider_OnMessagePublish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnMessageDelivered(ctx context.Context, in *MessageDeliveredRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnMessageDelivered_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnMessageDropped(ctx context.Context, in *MessageDroppedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnMessageDropped_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookProviderClient) OnMessageAcked(ctx context.Context, in *MessageAckedRequest, opts ...grpc.CallOption) (*EmptySuccess, error) {
	out := new(EmptySuccess)
	err := c.cc.Invoke(ctx, HookProvider_OnMessageAcked_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookProviderServer is the server API for HookProvider service.
// All implementations must embed UnimplementedHookProviderServer
// for forward compatibility
type HookProviderServer interface {
	OnProviderLoaded(context.Context, *ProviderLoadedRequest) (*LoadedResponse, error)
	OnProviderUnloaded(context.Context, *ProviderUnloadedRequest) (*EmptySuccess, error)
	OnClientConnect(context.Context, *ClientConnectRequest) (*EmptySuccess, error)
	OnClientConnack(context.Context, *ClientConnackRequest) (*EmptySuccess, error)
	OnClientConnected(context.Context, *ClientConnectedRequest) (*EmptySuccess, error)
	OnClientDisconnected(context.Context, *ClientDisconnectedRequest) (*EmptySuccess, error)
	OnClientAuthenticate(context.Context, *ClientAuthenticateRequest) (*ValuedResponse, error)
	OnClientCheckAcl(context.Context, *ClientCheckAclRequest) (*ValuedResponse, error)
	OnClientSubscribe(context.Context, *ClientSubscribeRequest) (*EmptySuccess, error)
	OnClientUnsubscribe(context.Context, *ClientUnsubscribeRequest) (*EmptySuccess, error)
	OnSessionCreated(context.Context, *SessionCreatedRequest) (*EmptySuccess, error)
	OnSessionSubscribed(context.Context, *SessionSubscribedRequest) (*EmptySuccess, error)
	OnSessionUnsubscribed(context.Context, *SessionUnsubscribedRequest) (*EmptySuccess, error)
	OnSessionResumed(context.Context, *SessionResumedRequest) (*EmptySuccess, error)
	OnSessionDiscarded(context.Context, *SessionDiscardedRequest) (*EmptySuccess, error)
	OnSessionTakeovered(context.Context, *SessionTakeoveredRequest) (*EmptySuccess, error)
	OnSessionTerminated(context.Context, *SessionTerminatedRequest) (*EmptySuccess, error)
	OnMessagePublish(context.Context, *MessagePublishRequest) (*ValuedResponse, error)
	OnMessageDelivered(context.Context, *MessageDeliveredRequest) (*EmptySuccess, error)
	OnMessageDropped(context.Context, *MessageDroppedRequest) (*EmptySuccess, error)
	OnMessageAcked(context.Context, *MessageAckedRequest) (*EmptySuccess, error)
	mustEmbedUnimplementedHookProviderServer()
}

// UnimplementedHookProviderServer must be embedded to have forward compatible implementations.
type UnimplementedHookProviderServer struct {
}

func (UnimplementedHookProviderServer) OnProviderLoaded(context.Context, *ProviderLoadedRequest) (*LoadedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnProviderLoaded not implemented")
}
func (UnimplementedHookProviderServer) OnProviderUnloaded(context.Context, *ProviderUnloadedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnProviderUnloaded not implemented")
}
func (UnimplementedHookProviderServer) OnClientConnect(context.Context, *ClientConnectRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientConnect not implemented")
}
func (UnimplementedHookProviderServer) OnClientConnack(context.Context, *ClientConnackRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientConnack not implemented")
}
func (UnimplementedHookProviderServer) OnClientConnected(context.Context, *ClientConnectedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientConnected not implemented")
}
func (UnimplementedHookProviderServer) OnClientDisconnected(context.Context, *ClientDisconnectedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientDisconnected not implemented")
}
func (UnimplementedHookProviderServer) OnClientAuthenticate(context.Context, *ClientAuthenticateRequest) (*ValuedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientAuthenticate not implemented")
}
func (UnimplementedHookProviderServer) OnClientCheckAcl(context.Context, *ClientCheckAclRequest) (*ValuedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientCheckAcl not implemented")
}
func (UnimplementedHookProviderServer) OnClientSubscribe(context.Context, *ClientSubscribeRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientSubscribe not implemented")
}
func (UnimplementedHookProviderServer) OnClientUnsubscribe(context.Context, *ClientUnsubscribeRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClientUnsubscribe not implemented")
}
func (UnimplementedHookProviderServer) OnSessionCreated(context.Context, *SessionCreatedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionCreated not implemented")
}
func (UnimplementedHookProviderServer) OnSessionSubscribed(context.Context, *SessionSubscribedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionSubscribed not implemented")
}
func (UnimplementedHookProviderServer) OnSessionUnsubscribed(context.Context, *SessionUnsubscribedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionUnsubscribed not implemented")
}
func (UnimplementedHookProviderServer) OnSessionResumed(context.Context, *SessionResumedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionResumed not implemented")
}
func (UnimplementedHookProviderServer) OnSessionDiscarded(context.Context, *SessionDiscardedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionDiscarded not implemented")
}
func (UnimplementedHookProviderServer) OnSessionTakeovered(context.Context, *SessionTakeoveredRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionTakeovered not implemented")
}
func (UnimplementedHookProviderServer) OnSessionTerminated(context.Context, *SessionTerminatedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSessionTerminated not implemented")
}
func (UnimplementedHookProviderServer) OnMessagePublish(context.Context, *MessagePublishRequest) (*ValuedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessagePublish not implemented")
}
func (UnimplementedHookProviderServer) OnMessageDelivered(context.Context, *MessageDeliveredRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessageDelivered not implemented")
}
func (UnimplementedHookProviderServer) OnMessageDropped(context.Context, *MessageDroppedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessageDropped not implemented")
}
func (UnimplementedHookProviderServer) OnMessageAcked(context.Context, *MessageAckedRequest) (*EmptySuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessageAcked not implemented")
}
func (UnimplementedHookProviderServer) mustEmbedUnimplementedHookProviderServer() {}

// UnsafeHookProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HookProviderServer will
// result in compilation errors.
type UnsafeHookProviderServer interface {
	mustEmbedUnimplementedHookProviderServer()
}

func RegisterHookProviderServer(s grpc.ServiceRegistrar, srv HookProviderServer) {
	s.RegisterService(&HookProvider_ServiceDesc, srv)
}

func _HookProvider_OnProviderLoaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderLoadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnProviderLoaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnProviderLoaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnProviderLoaded(ctx, req.(*ProviderLoadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnProviderUnloaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderUnloadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnProviderUnloaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnProviderUnloaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnProviderUnloaded(ctx, req.(*ProviderUnloadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientConnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientConnect(ctx, req.(*ClientConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientConnack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConnackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientConnack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientConnack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientConnack(ctx, req.(*ClientConnackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientConnected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConnectedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientConnected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientConnected_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientConnected(ctx, req.(*ClientConnectedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientDisconnected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientDisconnectedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientDisconnected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientDisconnected_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientDisconnected(ctx, req.(*ClientDisconnectedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientAuthenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientAuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientAuthenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientAuthenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientAuthenticate(ctx, req.(*ClientAuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientCheckAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientCheckAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientCheckAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientCheckAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientCheckAcl(ctx, req.(*ClientCheckAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientSubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientSubscribe(ctx, req.(*ClientSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnClientUnsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientUnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnClientUnsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnClientUnsubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnClientUnsubscribe(ctx, req.(*ClientUnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionCreatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionCreated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionCreated(ctx, req.(*SessionCreatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionSubscribed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionSubscribedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionSubscribed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionSubscribed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionSubscribed(ctx, req.(*SessionSubscribedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionUnsubscribed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionUnsubscribedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionUnsubscribed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionUnsubscribed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionUnsubscribed(ctx, req.(*SessionUnsubscribedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionResumed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionResumedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionResumed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionResumed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionResumed(ctx, req.(*SessionResumedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionDiscarded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionDiscardedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionDiscarded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionDiscarded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionDiscarded(ctx, req.(*SessionDiscardedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionTakeovered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionTakeoveredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionTakeovered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionTakeovered_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionTakeovered(ctx, req.(*SessionTakeoveredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnSessionTerminated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionTerminatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnSessionTerminated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnSessionTerminated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnSessionTerminated(ctx, req.(*SessionTerminatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnMessagePublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagePublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnMessagePublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnMessagePublish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnMessagePublish(ctx, req.(*MessagePublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnMessageDelivered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageDeliveredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnMessageDelivered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnMessageDelivered_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnMessageDelivered(ctx, req.(*MessageDeliveredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnMessageDropped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageDroppedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnMessageDropped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnMessageDropped_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnMessageDropped(ctx, req.(*MessageDroppedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookProvider_OnMessageAcked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageAckedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookProviderServer).OnMessageAcked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HookProvider_OnMessageAcked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookProviderServer).OnMessageAcked(ctx, req.(*MessageAckedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HookProvider_ServiceDesc is the grpc.ServiceDesc for HookProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HookProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emqx.exhook.v1.HookProvider",
	HandlerType: (*HookProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnProviderLoaded",
			Handler:    _HookProvider_OnProviderLoaded_Handler,
		},
		{
			MethodName: "OnProviderUnloaded",
			Handler:    _HookProvider_OnProviderUnloaded_Handler,
		},
		{
			MethodName: "OnClientConnect",
			Handler:    _HookProvider_OnClientConnect_Handler,
		},
		{
			MethodName: "OnClientConnack",
			Handler:    _HookProvider_OnClientConnack_Handler,
		},
		{
			MethodName: "OnClientConnected",
			Handler:    _HookProvider_OnClientConnected_Handler,
		},
		{
			MethodName: "OnClientDisconnected",
			Handler:    _HookProvider_OnClientDisconnected_Handler,
		},
		{
			MethodName: "OnClientAuthenticate",
			Handler:    _HookProvider_OnClientAuthenticate_Handler,
		},
		{
			MethodName: "OnClientCheckAcl",
			Handler:    _HookProvider_OnClientCheckAcl_Handler,
		},
		{
			MethodName: "OnClientSubscribe",
			Handler:    _HookProvider_OnClientSubscribe_Handler,
		},
		{
			MethodName: "OnClientUnsubscribe",
			Handler:    _HookProvider_OnClientUnsubscribe_Handler,
		},
		{
			MethodName: "OnSessionCreated",
			Handler:    _HookProvider_OnSessionCreated_Handler,
		},
		{
			MethodName: "OnSessionSubscribed",
			Handler:    _HookProvider_OnSessionSubscribed_Handler,
		},
		{
			MethodName: "OnSessionUnsubscribed",
			Handler:    _HookProvider_OnSessionUnsubscribed_Handler,
		},
		{
			MethodName: "OnSessionResumed",
			Handler:    _HookProvider_OnSessionResumed_Handler,
		},
		{
			MethodName: "OnSessionDiscarded",
			Handler:    _HookProvider_OnSessionDiscarded_Handler,
		},
		{
			MethodName: "OnSessionTakeovered",
			Handler:    _HookProvider_OnSessionTakeovered_Handler,
		},
		{
			MethodName: "OnSessionTerminated",
			Handler:    _HookProvider_OnSessionTerminated_Handler,
		},
		{
			MethodName: "OnMessagePublish",
			Handler:    _HookProvider_OnMessagePublish_Handler,
		},
		{
			MethodName: "OnMessageDelivered",
			Handler:    _HookProvider_OnMessageDelivered_Handler,
		},
		{
			MethodName: "OnMessageDropped",
			Handler:    _HookProvider_OnMessageDropped_Handler,
		},
		{
			MethodName: "OnMessageAcked",
			Handler:    _HookProvider_OnMessageAcked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/exhook.proto",
}
