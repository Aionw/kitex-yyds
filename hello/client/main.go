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
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kitex-yyds/kitex-yyds/consul"
	"github.com/kitex-yyds/kitex-yyds/kitex_gen/echo"
	"github.com/kitex-yyds/kitex-yyds/kitex_gen/echo/echoservice"
	"github.com/kitex-yyds/kitex-yyds/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
	tclient "github.com/kitex-yyds/kitex-yyds/tracer/client"
)

func main() {
	tracerOpt, closer := tclient.InitJaeger("hello-client")
	defer closer.Close()
	consul.Init()

	cli, err := echoservice.NewClient("hello",
		client.WithTransportProtocol(transport.GRPC),
		tracerOpt,
		client.WithMiddlewareBuilder(middleware.LogMiddlewareBuilder),
		client.WithResolver(&consul.ConsulResolver{
			ConsulClient: consul.Client,
		}))
	if err != nil {
		panic(err)
	}
	cliStream, err := cli.ClientSideStreaming(context.Background())
	if err != nil {
		panic(err)
	}
	for {
		req := &echo.Request{Msg: "hello"}
		if err := cliStream.Send(req); err != nil {
			panic(err)
		}
		fmt.Printf("request: %v\n", req)
		resp := &echo.Response{}
		err = cliStream.RecvMsg(resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("response: %v\n", resp)
		time.Sleep(time.Second)
	}
}
