// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"log"
	"time"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "grpc-opencensus/proto"
	"google.golang.org/grpc"
	// 声明为了启用监控需要添加的外部库。
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/examples/exporter"
)

const (
	address = "localhost:50051"
)

func main() {
	// Register stats and trace exporters to export
	// the collected data.
	// 注册统计数据和跟踪的导出器，以导出收集的数据。这里添加了PrintExporter, 它会将导出数据以日志的形式打印到控制台上。
	// 这只是为了展示功能，正常情况下，不推荐日志记录所有的生产环境负载。
	view.RegisterExporter(&exporter.PrintExporter{})

    // Register the view to collect gRPC client stats.
    // 注册视图以收集服务器请求的数量。这些是预定义的默认服务视图，会收集每个RPC所接收到的字节、每个RPC发送的字节、每个RPC的延迟以及完成的RPC。
    // 可以编写自己的视图来收集数据。
    if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
        log.Fatal(err)
    }

	// Set up a connection to the server.
	// 使用客户端统计数据的handler 建立到服务器端的连接。
	conn, err := grpc.Dial(address,
			grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	        grpc.WithInsecure(),
	        )
	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}
	// 在所有的事情完成后关闭连接。
	defer conn.Close()

	// 使用服务器端连接创建客户端存根。
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
