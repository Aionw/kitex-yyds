package main

import (
	api "github.com/kitex-yyds/kitex-yyds/hello_thrift/kitex_gen/api/hellothrift"
	"log"
)

func main() {
	svr := api.NewServer(new(HelloThriftImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
