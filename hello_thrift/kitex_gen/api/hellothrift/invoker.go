// Code generated by Kitex v0.0.3. DO NOT EDIT.

package hellothrift

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler api.HelloThrift, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}