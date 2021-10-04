// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_info product_info.proto
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	pb "token-based-authentication/client/ecommerce"
)

var (
	address = "localhost:50051"
	hostname = "localhost"
	crtFile = filepath.Join("src", "token-based-authentication", "certs", "server.crt")
)

func main() {
	// Set up the credentials for the connection.
	// 设置连接的凭证，需要提供OAuth令牌值来创建凭证。这里使用一个硬编码的字符串值作为令牌的值。
	perRPC := oauth.NewOauthAccess(fetchToken())
	// 调用oauth.NewJWTAccessFromFile初始化credentials.PerRPCCredentials，需要提供一个有效的令牌文件来创建凭证。
	//jwtCreds, err := oauth.NewJWTAccessFromFile("token.json")
	//if err != nil {
	//	log.Fatalf("failed to create JWT credentials: %v", err)
	//}
	// 调用oauth.NewServiceAccountFromFile来初始化credentials.PerRPCCredentials，需要提供一个有效的令牌文件来创建凭证。
	//googleCreds, err := oauth.NewServiceAccountFromFile("service-account.json", scope)
	//if err != nil {
	//	log.Fatalf("failed to create JWT credentials: %v", err)
	//}

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		// 配置gRPC DialOption，为同一个连接的所有RPC使用同一个令牌。
		// 如果想为每个调用使用专门的OAuth令牌，那么需要使用CallOption配置gRPC调用。
		grpc.WithPerRPCCredentials(perRPC),
		// 使用DialOption WithPerRPCCredentials配置gRPC dial, 为相同连接的所有RPC使用同一个JWT令牌。
		//grpc.WithPerRPCCredentials(jwtCreds),
		// 使用DialOption WithPerRPCCredentials配置gRPC dial，从而将认证令牌作为元数据应用于相同连接的所有RPC.
		//grpc.WithPerRPCCredentials(googleCreds),
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
	log.Printf("Product: ", product.String())
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
