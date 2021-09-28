package main

import (
	"context"
	pb "deadlines/order-service/order-service-gen"
	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const (
	port           = ":50051"
	orderBatchSize = 3
)

var orderMap = make(map[string]pb.Order)

type server struct {
	orderMap map[string]*pb.Order
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrappers.StringValue, error) {
	orderMap[orderReq.Id] = *orderReq

	sleepDuration  := 5
	log.Println("Sleeping for :",  sleepDuration, "s")

	time.Sleep(time.Duration(sleepDuration) * time.Second)

	//
	// 判断客户端是否已经满足超出截止时间的状态，随后就可以在服务器端废弃该RPC并返回一个错误。
	// 在Go语言中，这通常会通过非阻塞的select构造来实现。
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("RPC has reached deadline exceeded state : %s", ctx.Err())
		return nil, ctx.Err()
	}

	log.Println("Order : ",  orderReq.Id, " -> Added")
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 30.00}
}
