// Copyright 2021 CloudWeGo Authors
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
//

package main

import (
	"context"
	"fmt"
	"github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-yyds/kitex-yyds/consul"
	"github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api/hellothrift"
	"github.com/kitex-yyds/kitex-yyds/kitex_gen/echo"
	tclient "github.com/kitex-yyds/kitex-yyds/tracer/client"
)

type handler struct{}

var Client hellothrift.Client

func Init() {
	var err error
	tracerOpt, closer := tclient.InitJaeger("hello-thrift-client")
	defer closer.Close()
	Client, err = hellothrift.NewClient(
		"hello-thrift",
		tracerOpt,
		client.WithResolver(&consul.ConsulResolver{
			ConsulClient: consul.Client,
		}),
	)
	if err != nil {
		panic(err)
	}
}

func (handler) ClientSideStreaming(stream echo.EchoService_ClientSideStreamingServer) (err error) {
	_, err = Client.Echo(context.Background(), &api.Request{Message: "hello streaming"})
	if err != nil {
		return err
	}
	for {
		_, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println("client recv")
		resp := &echo.Response{Msg: "world"}
		if err := stream.SendMsg(resp); err != nil {
			return err
		}
		fmt.Println("client send")
	}
}

func (handler) ServerSideStreaming(req *echo.Request, stream echo.EchoService_ServerSideStreamingServer) (err error) {
	_, err = Client.Echo(context.Background(), &api.Request{Message: "hello streaming2"})
	_ = req
	for {
		resp := &echo.Response{Msg: "world"}
		if err := stream.Send(resp); err != nil {
			return err
		}
		fmt.Println("server recv")
		req := &echo.Request{Msg: "world"}
		if err := stream.SendMsg(req); err != nil {
			return err
		}
		fmt.Println("server send")
	}
}

func (handler) BidiSideStreaming(stream echo.EchoService_BidiSideStreamingServer) (err error) {
	_, err = Client.Echo(context.Background(), &api.Request{Message: "hello streaming3"})
	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				return
			}
		}
	}()
	for {
		resp := &echo.Response{Msg: "world"}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
