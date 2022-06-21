// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: nodeum/plugins/v1/finalizer.proto

package pluginsv1

import (
	fmt "fmt"
	_ "github.com/nodeum-io/nodeum-plugins/nodeum/common/v1"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FinalizerPluginService service

func NewFinalizerPluginServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FinalizerPluginService service

type FinalizerPluginService interface {
	// After a result has been received and before it's being handled by finalizer.
	// Progress results are excluded.
	AfterResultReceived(ctx context.Context, in *AfterResultReceivedRequest, opts ...client.CallOption) (*AfterResultReceivedResponse, error)
	// After the task has been executed
	AfterTaskExecuted(ctx context.Context, in *AfterTaskExecutedRequest, opts ...client.CallOption) (*AfterTaskExecutedResponse, error)
}

type finalizerPluginService struct {
	c    client.Client
	name string
}

func NewFinalizerPluginService(name string, c client.Client) FinalizerPluginService {
	return &finalizerPluginService{
		c:    c,
		name: name,
	}
}

func (c *finalizerPluginService) AfterResultReceived(ctx context.Context, in *AfterResultReceivedRequest, opts ...client.CallOption) (*AfterResultReceivedResponse, error) {
	req := c.c.NewRequest(c.name, "FinalizerPluginService.AfterResultReceived", in)
	out := new(AfterResultReceivedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *finalizerPluginService) AfterTaskExecuted(ctx context.Context, in *AfterTaskExecutedRequest, opts ...client.CallOption) (*AfterTaskExecutedResponse, error) {
	req := c.c.NewRequest(c.name, "FinalizerPluginService.AfterTaskExecuted", in)
	out := new(AfterTaskExecutedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FinalizerPluginService service

type FinalizerPluginServiceHandler interface {
	// After a result has been received and before it's being handled by finalizer.
	// Progress results are excluded.
	AfterResultReceived(context.Context, *AfterResultReceivedRequest, *AfterResultReceivedResponse) error
	// After the task has been executed
	AfterTaskExecuted(context.Context, *AfterTaskExecutedRequest, *AfterTaskExecutedResponse) error
}

func RegisterFinalizerPluginServiceHandler(s server.Server, hdlr FinalizerPluginServiceHandler, opts ...server.HandlerOption) error {
	type finalizerPluginService interface {
		AfterResultReceived(ctx context.Context, in *AfterResultReceivedRequest, out *AfterResultReceivedResponse) error
		AfterTaskExecuted(ctx context.Context, in *AfterTaskExecutedRequest, out *AfterTaskExecutedResponse) error
	}
	type FinalizerPluginService struct {
		finalizerPluginService
	}
	h := &finalizerPluginServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FinalizerPluginService{h}, opts...))
}

type finalizerPluginServiceHandler struct {
	FinalizerPluginServiceHandler
}

func (h *finalizerPluginServiceHandler) AfterResultReceived(ctx context.Context, in *AfterResultReceivedRequest, out *AfterResultReceivedResponse) error {
	return h.FinalizerPluginServiceHandler.AfterResultReceived(ctx, in, out)
}

func (h *finalizerPluginServiceHandler) AfterTaskExecuted(ctx context.Context, in *AfterTaskExecutedRequest, out *AfterTaskExecutedResponse) error {
	return h.FinalizerPluginServiceHandler.AfterTaskExecuted(ctx, in, out)
}
