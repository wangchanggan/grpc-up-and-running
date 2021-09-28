package main

import (
	"context"
	pb "error-handling/order-service/order-service-gen"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Add Order
	// This is an invalid order
	// 这是一个非法订单。
	order1 := pb.Order{Id: "-1", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	// 调用 AddOrder远程方法并将错误赋值给addOrderError。
	res, addOrderError := client.AddOrder(ctx, &order1)

	if addOrderError != nil {
		// 使用status包获取错误码。
		errorCode := status.Code(addOrderError)
		// 检查InvalidArgument 错误码。
		if errorCode == codes.InvalidArgument {
			log.Printf("Invalid Argument Error : %s", errorCode)
			// 从错误中获取错误状态。
			errorStatus := status.Convert(addOrderError)
			for _, d := range errorStatus.Details() {
				switch info := d.(type) {
				// 检查BadRequest_FieldViolation错误类型。
				case *epb.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s", info)
				default:
					log.Printf("Unexpected error type: %s", info)
				}
			}
		} else {
			log.Printf("Unhandled error : %s ", errorCode)
		}
	} else {
		log.Print("AddOrder Response -> ", res.Value)
	}
}
