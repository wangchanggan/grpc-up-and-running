// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	pb "grpc-prometheus/proto"
	"google.golang.org/grpc"
	// 声明要启用监控功能所需的外部库。gRPC提供了预定义的一组拦截器以支持Prometheus监控。在这里将使用这些拦截器。
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	// 借助添加的自定义度量指标计数器，能够为监控添加更多的度量指标。
	// 假设想收集相同名称的商品向商品管理系统中添加的次数，可以在AddProduct方法中添加名为customizedCounterMetric的新度量指标。
	customizedCounterMetric.WithLabelValues(in.Name).Inc()
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

var (
	// Create a metrics registry.
	// 创建度量指标的注册中心。它会持有系统中所有注册的数据收集器。如需添加新的收集器，就要在这个注册中心中对其进行注册。
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	// 创建标准的服务器端度量指标。这是在库中预先定义好的度量指标。
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// 创建名为product_mgt_server_handle_count的自定义度量指标计数器。
    customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
        Name: "product_mgt_server_handle_count",
        Help: "Total number of RPCs handled on the server.",
    }, []string{"name"})
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	// 将标准的服务器度量指标和自定义的度量指标收集器注册到第2步所创建的注册中心里。
	reg.MustRegister(grpcMetrics, customizedCounterMetric)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a HTTP server for prometheus.
	// 为Prometheus创建HTTP服务器。在端口9092上以上下文/metrics 开头的HTTP端点用来进行度量指标收集。
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092)}

	// Create a gRPC Server with gRPC interceptor.
	// 使用度量指标栏截器创建gRPC服务器。这里使用了grpcMetrics.StreamServerInterceptor，因为具有一元服务。
	// 还有另一个适用于流服务的拦截器，名为grpcMetrics.UnaryServerInterceptor
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	pb.RegisterProductInfoServer(grpcServer, &server{})
    // Initialize all metrics.
	// 初始化所有的标准度量指标。
    grpcMetrics.InitializeMetrics(grpcServer)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
