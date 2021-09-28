package main

import (
	"context"
	pb "deadlines/order-service/order-service-gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// Setting up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)
	// 在当前上下文中设置2秒的截止时间。
	clientDeadline := time.Now().Add(time.Duration(2 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	// Add Order
	order1 := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	// 调用AddOrder远程方法并将可能出现的错误捕获到addErr中。
	res, addErr := client.AddOrder(ctx, &order1)

	if addErr != nil {
		// 使用status包以确定错误码。
		got := status.Code(addErr)
		// 如果调用超出了指定的截止时间，它应该返回DEADLINE_EXCEEDED类型的错误。
		log.Printf("Error Occured -> addOrder : , %v:", got)
	} else {
		log.Print("AddOrder Response -> ", res.Value)
	}
}
