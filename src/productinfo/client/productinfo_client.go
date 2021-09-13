package main

import (
	"context"
	"log"
	"time"
    // 导入protobuf编译器生成代码所在的包
	pb "productinfo/client/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	// 根据提供的地址(localhost: 50051) 创建到服务器端的连接。这里创建了一个客户端和服务器端之间的连接，但它目前不安全。
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 所有事情都完成后，关闭连接。
	defer conn.Close()

	// 传递连接并创建存根文件。这个实例包含可调用服务器的所有远程方法。
	c := pb.NewProductInfoClient(conn)

	// Contact the server and print out its response.
	name := "Apple iPhone 11"
	description := "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode."
	price := float32(699.00)

	// 创建Context以传递给远程调用。这里的Context对象包含一些元数据， 如终端用户的标识、授权令牌以及请求的截止时间，该对象会在请求的生命周期内一直存在。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 使用商品的详情信息调用AddProduct方法。如果操作成功完成，就会返回-一个商品ID，否则将返回一个错误。
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	// 使用商品ID来调用GetProduct方法。如果操作成功完成，将返回商品详情，否则会返回一个错误。
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %v", product.String())
}
