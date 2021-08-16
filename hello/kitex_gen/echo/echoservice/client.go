// Code generated by Kitex v0.0.3. DO NOT EDIT.

package echoservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/streaming"
	_ "github.com/cloudwego/kitex/transport"
	"github.com/kitex-yyds/kitex-yyds/hello/kitex_gen/echo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ClientSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream EchoService_ClientSideStreamingClient, err error)
	ServerSideStreaming(ctx context.Context, Req *echo.Request, callOptions ...callopt.Option) (stream EchoService_ServerSideStreamingClient, err error)
	BidiSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream EchoService_BidiSideStreamingClient, err error)
}

type EchoService_ClientSideStreamingClient interface {
	streaming.Stream
	Send(*echo.Request) error
	CloseAndRecv() (*echo.Response, error)
}

type EchoService_ServerSideStreamingClient interface {
	streaming.Stream
	Recv() (*echo.Response, error)
}

type EchoService_BidiSideStreamingClient interface {
	streaming.Stream
	Send(*echo.Request) error
	Recv() (*echo.Response, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoServiceClient struct {
	*kClient
}

func (p *kEchoServiceClient) ClientSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream EchoService_ClientSideStreamingClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ClientSideStreaming(ctx)
}

func (p *kEchoServiceClient) ServerSideStreaming(ctx context.Context, Req *echo.Request, callOptions ...callopt.Option) (stream EchoService_ServerSideStreamingClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ServerSideStreaming(ctx, Req)
}

func (p *kEchoServiceClient) BidiSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream EchoService_BidiSideStreamingClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BidiSideStreaming(ctx)
}
