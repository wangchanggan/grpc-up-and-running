package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	// 导入生成的反向代理代码所在的包。
	gw "grpc-gateway/proto"
)

var (
	// command-line options:
	// gRPC server endpoint
	// 声明gRPC服务器端点URL，确保后端gRPC服务器在所述的端点上正常运行。
	grpcServerEndpoint = "localhost:50051"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 使用代理handler注册gRPC服务器端点。在运行时，请求多路转换器(multiplexer)将HTTP请求匹配为模式，并调用对应的handler。
	err := gw.RegisterProductInfoHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC service endpoint: %v", err)
		return
	}
	// 开始在端口8081上监听HTTP请求。
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}
}
