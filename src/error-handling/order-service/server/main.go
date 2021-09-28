package main

import (
	"context"
	pb "error-handling/order-service/order-service-gen"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
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
	// 非法请求，需要生成一个错误并将其返回给客户端。
	// 包含错误类型BadRequest_FieldViolation的所有错误详情。

	if orderReq.Id == "-1" {
		log.Printf("Order ID is invalid! -> Received Order ID %s", orderReq.Id)

		// 创建一个错误码为InvalidArgument的新错误状态。
		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		ds, err := errorStatus.WithDetails(
			&epb.BadRequest_FieldViolation{
				Field:       "ID",
				Description: fmt.Sprintf("Order ID received is not valid %s : %s", orderReq.Id, orderReq.Description),
			},
		)
		if err != nil {
			return nil, errorStatus.Err()
		}

		// 返回生成的错误。
		return nil, ds.Err()
	} else {
		orderMap[orderReq.Id] = *orderReq
		log.Println("Order : ", orderReq.Id, " -> Added")
		return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
	}
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
