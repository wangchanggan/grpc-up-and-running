// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"net/http"
    // 声明要启用监控功能所需要的外部库。
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "grpc-prometheus/proto"
	"google.golang.org/grpc"
	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	address = "localhost:50051"
)

func main() {
	// Create a metrics registry.
	// 创建度量指标的注册中心。与服务器端代码类似，它会持有系统中所有注册的数据收集器。
	// 如果要添加新的收集器，就需要在这个注册中心中对共进行注册。
	reg := prometheus.NewRegistry()
	// Create some standard client metrics.
	// 创建标准的客户端度量指标，这是在库中预定义好的度量指标。
	grpcMetrics := grpc_prometheus.NewClientMetrics()
	// Register client metrics to registry.
	// 注册标准的客户端度量指标到所创建的注册中心里。
	reg.MustRegister(grpcMetrics)

	// Set up a connection to the server.
	// 使用度量指标拦截器创建到gRPC服务器端的连接。这里使用了grpcMetrics.UnaryClientInterceptor，因为具有一元客户端。
	// 还有另一个适用于流客户端的拦截器，叫作grpcMetrics.StreamClientInterceptor.
	conn, err := grpc.Dial(address,
			grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
	        grpc.WithInsecure(),
	        )
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

    // Create a HTTP server for prometheus.
	// 为Prometheus创建HTTP服务器。在9094端口上以上下文/metrics开头的HTTP端点用来进行度量指标收集。
    httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9094)}

    // Start your http server for prometheus.
    go func() {
        if err := httpServer.ListenAndServe(); err != nil {
            log.Fatal("Unable to start a http server.")
        }
    }()

	c := pb.NewProductInfoClient(conn)

	for {
        // Contact the server and print out its response.
        name := "Sumsung S10"
        description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
        price := float32(700.0)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
        if err != nil {
            log.Fatalf("Could not add product: %v", err)
        }
        log.Printf("Product ID: %s added successfully", r.Value)

        product, err := c.GetProduct(ctx, &wrapper.StringValue{Value: r.Value})
        if err != nil {
            log.Fatalf("Could not get product: %v", err)
        }
        log.Printf("Product: %s", product.String())
        time.Sleep(3 * time.Second)
    }
}
