// Code generated by Kitex v0.0.3. DO NOT EDIT.

package echoservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/kitex-yyds/kitex-yyds/kitex_gen/echo"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return echoServiceServiceInfo
}

var echoServiceServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "EchoService"
	handlerType := (*echo.EchoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ClientSideStreaming": kitex.NewMethodInfo(clientSideStreamingHandler, newClientSideStreamingArgs, newClientSideStreamingResult, false),
		"ServerSideStreaming": kitex.NewMethodInfo(serverSideStreamingHandler, newServerSideStreamingArgs, newServerSideStreamingResult, false),
		"BidiSideStreaming":   kitex.NewMethodInfo(bidiSideStreamingHandler, newBidiSideStreamingArgs, newBidiSideStreamingResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "echo",
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.0.3",
		Extra:           extra,
	}
	return svcInfo
}

func clientSideStreamingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &echoServiceClientSideStreamingServer{st}
	return handler.(echo.EchoService).ClientSideStreaming(stream)
}

type echoServiceClientSideStreamingClient struct {
	streaming.Stream
}

func (x *echoServiceClientSideStreamingClient) Send(m *echo.Request) error {
	return x.Stream.SendMsg(m)
}
func (x *echoServiceClientSideStreamingClient) CloseAndRecv() (*echo.Response, error) {
	if err := x.Stream.Close(); err != nil {
		return nil, err
	}
	m := new(echo.Response)
	return m, x.Stream.RecvMsg(m)
}

type echoServiceClientSideStreamingServer struct {
	streaming.Stream
}

func (x *echoServiceClientSideStreamingServer) SendAndClose(m *echo.Response) error {
	return x.Stream.SendMsg(m)
}

func (x *echoServiceClientSideStreamingServer) Recv() (*echo.Request, error) {
	m := new(echo.Request)
	return m, x.Stream.RecvMsg(m)
}

func newClientSideStreamingArgs() interface{} {
	return &ClientSideStreamingArgs{}
}

func newClientSideStreamingResult() interface{} {
	return &ClientSideStreamingResult{}
}

type ClientSideStreamingArgs struct {
	Req *echo.Request
}

func (p *ClientSideStreamingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ClientSideStreamingArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ClientSideStreamingArgs) Unmarshal(in []byte) error {
	msg := new(echo.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ClientSideStreamingArgs_Req_DEFAULT *echo.Request

func (p *ClientSideStreamingArgs) GetReq() *echo.Request {
	if !p.IsSetReq() {
		return ClientSideStreamingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ClientSideStreamingArgs) IsSetReq() bool {
	return p.Req != nil
}

type ClientSideStreamingResult struct {
	Success *echo.Response
}

var ClientSideStreamingResult_Success_DEFAULT *echo.Response

func (p *ClientSideStreamingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ClientSideStreamingResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ClientSideStreamingResult) Unmarshal(in []byte) error {
	msg := new(echo.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ClientSideStreamingResult) GetSuccess() *echo.Response {
	if !p.IsSetSuccess() {
		return ClientSideStreamingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ClientSideStreamingResult) SetSuccess(x interface{}) {
	p.Success = x.(*echo.Response)
}

func (p *ClientSideStreamingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func serverSideStreamingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &echoServiceServerSideStreamingServer{st}
	req := new(echo.Request)
	if err := st.RecvMsg(req); err != nil {
		return err
	}
	return handler.(echo.EchoService).ServerSideStreaming(req, stream)
}

type echoServiceServerSideStreamingClient struct {
	streaming.Stream
}

func (x *echoServiceServerSideStreamingClient) Recv() (*echo.Response, error) {
	m := new(echo.Response)
	return m, x.Stream.RecvMsg(m)
}

type echoServiceServerSideStreamingServer struct {
	streaming.Stream
}

func (x *echoServiceServerSideStreamingServer) Send(m *echo.Response) error {
	return x.Stream.SendMsg(m)
}

func newServerSideStreamingArgs() interface{} {
	return &ServerSideStreamingArgs{}
}

func newServerSideStreamingResult() interface{} {
	return &ServerSideStreamingResult{}
}

type ServerSideStreamingArgs struct {
	Req *echo.Request
}

func (p *ServerSideStreamingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ServerSideStreamingArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ServerSideStreamingArgs) Unmarshal(in []byte) error {
	msg := new(echo.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ServerSideStreamingArgs_Req_DEFAULT *echo.Request

func (p *ServerSideStreamingArgs) GetReq() *echo.Request {
	if !p.IsSetReq() {
		return ServerSideStreamingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ServerSideStreamingArgs) IsSetReq() bool {
	return p.Req != nil
}

type ServerSideStreamingResult struct {
	Success *echo.Response
}

var ServerSideStreamingResult_Success_DEFAULT *echo.Response

func (p *ServerSideStreamingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ServerSideStreamingResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ServerSideStreamingResult) Unmarshal(in []byte) error {
	msg := new(echo.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ServerSideStreamingResult) GetSuccess() *echo.Response {
	if !p.IsSetSuccess() {
		return ServerSideStreamingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ServerSideStreamingResult) SetSuccess(x interface{}) {
	p.Success = x.(*echo.Response)
}

func (p *ServerSideStreamingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func bidiSideStreamingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &echoServiceBidiSideStreamingServer{st}
	return handler.(echo.EchoService).BidiSideStreaming(stream)
}

type echoServiceBidiSideStreamingClient struct {
	streaming.Stream
}

func (x *echoServiceBidiSideStreamingClient) Send(m *echo.Request) error {
	return x.Stream.SendMsg(m)
}
func (x *echoServiceBidiSideStreamingClient) Recv() (*echo.Response, error) {
	m := new(echo.Response)
	return m, x.Stream.RecvMsg(m)
}

type echoServiceBidiSideStreamingServer struct {
	streaming.Stream
}

func (x *echoServiceBidiSideStreamingServer) Send(m *echo.Response) error {
	return x.Stream.SendMsg(m)
}

func (x *echoServiceBidiSideStreamingServer) Recv() (*echo.Request, error) {
	m := new(echo.Request)
	return m, x.Stream.RecvMsg(m)
}

func newBidiSideStreamingArgs() interface{} {
	return &BidiSideStreamingArgs{}
}

func newBidiSideStreamingResult() interface{} {
	return &BidiSideStreamingResult{}
}

type BidiSideStreamingArgs struct {
	Req *echo.Request
}

func (p *BidiSideStreamingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in BidiSideStreamingArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *BidiSideStreamingArgs) Unmarshal(in []byte) error {
	msg := new(echo.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var BidiSideStreamingArgs_Req_DEFAULT *echo.Request

func (p *BidiSideStreamingArgs) GetReq() *echo.Request {
	if !p.IsSetReq() {
		return BidiSideStreamingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BidiSideStreamingArgs) IsSetReq() bool {
	return p.Req != nil
}

type BidiSideStreamingResult struct {
	Success *echo.Response
}

var BidiSideStreamingResult_Success_DEFAULT *echo.Response

func (p *BidiSideStreamingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in BidiSideStreamingResult")
	}
	return proto.Marshal(p.Success)
}

func (p *BidiSideStreamingResult) Unmarshal(in []byte) error {
	msg := new(echo.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BidiSideStreamingResult) GetSuccess() *echo.Response {
	if !p.IsSetSuccess() {
		return BidiSideStreamingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BidiSideStreamingResult) SetSuccess(x interface{}) {
	p.Success = x.(*echo.Response)
}

func (p *BidiSideStreamingResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ClientSideStreaming(ctx context.Context) (EchoService_ClientSideStreamingClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "ClientSideStreaming", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &echoServiceClientSideStreamingClient{res.Stream}
	return stream, nil
}

func (p *kClient) ServerSideStreaming(ctx context.Context, req *echo.Request) (EchoService_ServerSideStreamingClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "ServerSideStreaming", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &echoServiceServerSideStreamingClient{res.Stream}
	if err := stream.Stream.SendMsg(req); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}

func (p *kClient) BidiSideStreaming(ctx context.Context) (EchoService_BidiSideStreamingClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "BidiSideStreaming", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &echoServiceBidiSideStreamingClient{res.Stream}
	return stream, nil
}
