// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.21.1
// source: chatproxy/v1/chatproxy.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationChatProxySayHello = "/chatproxy.v1.ChatProxy/SayHello"

type ChatProxyHTTPServer interface {
	// SayHello Sends a greeting
	SayHello(context.Context, *ChatRequest) (*ChatReply, error)
}

func RegisterChatProxyHTTPServer(s *http.Server, srv ChatProxyHTTPServer) {
	r := s.Route("/")
	r.POST("/chatproxy", _ChatProxy_SayHello0_HTTP_Handler(srv))
}

func _ChatProxy_SayHello0_HTTP_Handler(srv ChatProxyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatProxySayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*ChatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ChatReply)
		return ctx.Result(200, reply)
	}
}

type ChatProxyHTTPClient interface {
	SayHello(ctx context.Context, req *ChatRequest, opts ...http.CallOption) (rsp *ChatReply, err error)
}

type ChatProxyHTTPClientImpl struct {
	cc *http.Client
}

func NewChatProxyHTTPClient(client *http.Client) ChatProxyHTTPClient {
	return &ChatProxyHTTPClientImpl{client}
}

func (c *ChatProxyHTTPClientImpl) SayHello(ctx context.Context, in *ChatRequest, opts ...http.CallOption) (*ChatReply, error) {
	var out ChatReply
	pattern := "/chatproxy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatProxySayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
