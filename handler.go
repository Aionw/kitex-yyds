package main

import (
	"context"
	"github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api"
)

// HelloThriftImpl implements the last service interface defined in the IDL.
type HelloThriftImpl struct{}

// Echo implements the HelloThriftImpl interface.
func (s *HelloThriftImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
