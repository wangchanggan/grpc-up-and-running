// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_info product_info.proto
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc/credentials"
	"log"
	"path/filepath"
	"time"
	pb "basic-authentication/client/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(filepath.Join("src", "basic-authentication", "certs", "server.crt"),
		"localhost")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// 使用有效的用户凭证(用户名和密码)初始化auth变量。auth 变量存放了要使用的值。
	auth := basicAuth{
		username: "admin",
		password: "admin",
	}
	opts := []grpc.DialOption{
		// 传递auth变量给grpc.WithPerRPCCredentials函数。该函数接受一个接口作为参数。
		// 因为定义的认证结构符合该接口，所以可以传递变量。
		grpc.WithPerRPCCredentials(auth),
		// transport credentials.
		grpc.WithTransportCredentials(creds),
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
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

// 定义结构体来存放要注入RPC的字段集合(用户的凭证，如用户名和密码)。
type basicAuth struct {
	username string
	password string
}

// 实现GetRequestMetadata方法，并将用户凭证转换成请求元数据。
// 键是Authorization，值则由Basic和加上<用户名>:<密码>的base64算法计算结果所组成。
func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.username + ":" + b.password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

// 声明在传递凭证时是否需要启用通道安全性。
func (b basicAuth) RequireTransportSecurity() bool {
	return true
}
