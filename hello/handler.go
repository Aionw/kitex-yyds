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
	echo "github.com/cloudwego/kitex-examples/hello/kitex_gen/echo"
)

type handler struct{}

func (handler) ClientSideStreaming(stream echo.EchoService_ClientSideStreamingServer) (err error) {
	for {
		_, err := stream.Recv()
		if err != nil {
			return err
		}
	}
}

func (handler) ServerSideStreaming(req *echo.Request, stream echo.EchoService_ServerSideStreamingServer) (err error) {
	_ = req
	for {
		resp := &echo.Response{Msg: "world"}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func (handler) BidiSideStreaming(stream echo.EchoService_BidiSideStreamingServer) (err error) {
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
