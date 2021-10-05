// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	pb "grpc-opencensus/proto"
	"google.golang.org/grpc"
	// 为了启用监控，指明需要添加的外部库。gRPC OpenCensus提供了一组预先定义好的handler以支持OpenCensus监控。
	// 这里会使用这些handler.
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"
	"go.opencensus.io/examples/exporter"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &wrapper.StringValue{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, nil
	}
	return nil, errors.New("Product does not exist for the ID" + in.Value)
}

func main() {
	// Start z-Pages server.
	// 启动一台z-Pages服务器。在端口8081的/debug上下文中启动一个HTTP端点，实现度量指标的可视化。
	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe("127.0.0.1:8081", mux))
	}()

    // Register stats and trace exporters to export
    // the collected data.
	// 注册统计导出器以导出收集的数据。这里添加了PrintExporter，它会将导出数据以日志的形式打印到控制台上。
	// 这只是为了展示功能，正常情况下，不推荐日志记录所有的生产环境负载。
    view.RegisterExporter(&exporter.PrintExporter{})

	// Register the views to collect server request count.
	// 注册视图以收集服务器请求的数量。
	// 这些是预定义的默认服务视图，会收集每个RPC所接收的字节、每个RPC发送的字节、每个RPC的延迟以及完成的RPC。
	// 可以编写自己的视图来收集数据。
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	// Create a gRPC Server with stats handler.
	// 使用数据统计handler来创建gRPC服务器端。
	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	// 注册ProductInfo服务到服务器端上。
	pb.RegisterProductInfoServer(grpcServer, &server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 开始在端口(50051) 上监听传入的消息。
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
