package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	pb "metadata/order-service/order-service-gen"
	"net"
	"strings"
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
	log.Println("Order : ", orderReq.Id, " -> Added")

	// ***** Reading Metadata from Client *****
	// 一元RPC，从远程方法传入的上下文中读取元数据map
	md, metadataAvailable := metadata.FromIncomingContext(ctx)
	if !metadataAvailable {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["timestamp"]; ok {
		fmt.Printf("timestamp from metadata:\n")
		for i, e := range t {
			fmt.Printf("====> Metadata %d. %s\n", i, e)
		}
	}

	// Creating and sending a header.
	header := metadata.New(map[string]string{"location": "San Jose", "timestamp": time.Now().Format(time.StampNano)})
	trailer := metadata.New(map[string]string{"location": "San Jose", "timestamp": time.Now().Format(time.StampNano)})
	// 以头信息的形式发送元数据。
	grpc.SendHeader(ctx, header)
	// 和trailer一起发送元数据。
	grpc.SetTrailer(ctx, trailer)

	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Server-side Streaming RPC
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	// 流RPC，从流中获取上下文并从中读取元数据。
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Print("Metadata : ", md)
	} else {
		log.Print("failed to get metadata")
	}

	defer func() {
		trailer := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
		stream.SetTrailer(trailer)
	}()

	header := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(time.StampNano)})
	trailer := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(time.StampNano)})
	// 在流中，以头信息的形式发送元数据。
	stream.SendHeader(header)
	// 和流的trailer一起发送元数据。
	stream.SetTrailer(trailer)

	for key, order := range orderMap {
		for _, itemStr := range order.Items {
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				log.Print("Matching Order Found : "+key, " -> Writing Order to the stream ... ")
				stream.Send(&order)
				break
			}
		}
	}

	return nil
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
