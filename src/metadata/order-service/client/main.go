package main

import (
	"context"
	"fmt"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "metadata/order-service/order-service-gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
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

	// originalCtx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	// defer cancel()

	// ****** Metadata : Creation *****
	// 创建元数据
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"kn", "vn",
	)
	// 基于新的元数据创建新的上下文。
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)
	// 在现有的上下文中附加更多的元数据。
	ctxA := metadata.AppendToOutgoingContext(mdCtx, "k1", "v1", "k1", "v2", "k2", "v3")

	// RPC using the context with new metadata.
	// 用来存储RPC所返回的头信息和trailer的变量。
	var header, trailer metadata.MD

	// RPC: Add Order
	order1 := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	// 一元RPC使用带有元数据的新上下文。
	// 传递头信息和trailer引用来存储一元RPC所返回的值。
	res, _ := client.AddOrder(ctxA, &order1, grpc.Header(&header), grpc.Trailer(&trailer))

	log.Print("AddOrder Response -> ", res.Value)

	// Reading the headers
	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := header["location"]; ok {
		log.Printf("location from header:\n")
		for i, e := range l {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in header")
	}

	// Search Order
	// 相同的上下文也可用于流RPC。
	searchStream, _ := client.SearchOrders(ctxA, &wrapper.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			log.Print("Search Result : ", searchOrder)

			// 从流中获取头信息。
			streamHeader, err := searchStream.Header()
			if err != nil {
				log.Print("Search Header Error : ", err)
			} else {
				log.Print("Search Header : ", streamHeader)
			}

			// 从流中获取trailer, 用于发送状态码和状态消息。
			streamTrailer := searchStream.Trailer()
			if err != nil {
				log.Print("Search Trailer Error : ", err)
			} else {
				log.Print("Search Trailer : ", streamTrailer)
			}
		}
	}
}
