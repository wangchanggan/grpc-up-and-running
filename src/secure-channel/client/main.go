// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_info product_info.proto
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"
	"path/filepath"
	"time"
	pb "secure-channel/client/ecommerce"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
	hostname = "localhost"
	crtFile = filepath.Join("src", "secure-channel", "certs", "server.crt")
)

func main() {
	// 读取并解析公开证书，创建启用TLS的证书。
	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		// transport credentials.
		// 以DialOption的形式添加传输凭证。
		grpc.WithTransportCredentials(creds),
	}

	// Set up a connection to the server.
	// 通过传入dial选项，建立到服务器的安全连接。
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 所有事情完成后关闭连接。
	defer conn.Close()
	// 传入连接并创建存根。该存根实例包含了调用服务器的所有远程方法。
	c := pb.NewProductInfoClient(conn)

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

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s", product.String())
}
