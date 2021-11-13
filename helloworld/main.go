// Micro Srv

package main

import (
	"context"
	"github.com/asim/go-micro/v2"
	"log"

	// 引用上面生成的proto文件
	proto "micro/helloworld/greeter"
)

type Greeter struct{}

func (g *Greeter) Hello2(ctx context.Context, maskRequest *proto.HelloFieldMaskRequest, resp *proto.FastCodes) error {
	//resp.FieldMask = maskRequest.FieldMask.GetPaths()
	//fmt.Println("maskRequest.FieldMask.GetPaths() : ", maskRequest.TaskId)
	resp.FastCodes = make([]*proto.FastCode, 1)

	resp.FastCodes[0] = &proto.FastCode{
		FastCodeNum:   "num",
		FastCodeDesc:  "desc",
		FastCodeType:  "yes",
	}
	return nil
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

//go:generate protoc --proto_path=. --proto_path=C:/Users/11/go/bin/googleapis --go_out=. --micro_out=. greeter.proto
//go:generate protoc-go-inject-tag -input=./greeter/greeter.pb.go

func main() {
	// new一个微服务出来
	service := micro.NewService(
		//micro.Transport(grpc.NewTransport()),
		micro.Name("greeter1"),
		micro.Address(":38313"),
		micro.Version("latest"),
	)

	// 可选 解析命令行
	service.Init()

	// 注册 handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
