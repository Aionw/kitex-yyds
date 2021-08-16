module github.com/kitex-yyds/kitex-yyds

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/kitex v0.0.3
	github.com/cloudwego/kitex-examples v0.0.0-20210804074130-fefb675298e3
	github.com/kitex-contrib/tracer-opentracing v0.0.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	google.golang.org/protobuf v1.27.1
)
