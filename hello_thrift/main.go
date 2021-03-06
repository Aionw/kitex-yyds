package main

import (
	"log"

	"github.com/kitex-yyds/kitex-yyds/consul"
	api "github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api/hellothrift"
	tclient "github.com/kitex-yyds/kitex-yyds/tracer/server"
)

func main() {
	tracerOpt, closer := tclient.InitJaeger("hello-thrift-server")
	defer closer.Close()
	consul.Init()
	consul.Register("hello-thrift", 8888)
	svr := api.NewServer(new(HelloImpl), tracerOpt)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
