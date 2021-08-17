// Code generated by Kitex v0.0.3. DO NOT EDIT.
package hellothrift

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.HelloThrift, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
