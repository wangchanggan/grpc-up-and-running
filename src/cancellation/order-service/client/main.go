package main

import (
	pb "cancellation/order-service/order-service-gen"
	"context"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
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

	// 获取对cancel的引用。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	// Process Order
	// 调用流RPC。
	streamProcOrder, _ := client.ProcessOrders(ctx)
	// 通过流发送消息给服务。
	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "102"})
	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "103"})
	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "104"})

	channel := make(chan bool, 1)

	go asncClientBidirectionalRPC(streamProcOrder, channel)
	time.Sleep(time.Millisecond * 1000)

	// Cancelling the RPC
	// 在客户端，取消RPC (终止RPC)。
	cancel()
	// 当前上下文的状态。
	log.Printf("RPC Status : %s", ctx.Err())

	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "101"})
	_ = streamProcOrder.CloseSend()

	<-channel

}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan bool) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder != nil {
			// 当试图从已取消的上下文中接收消息时，会返回上下文已取消的错误。
			log.Printf("Error Receiving messages %v", errProcOrder)
			break
		} else {
			if errProcOrder == io.EOF {
				break
			}
			log.Printf("Combined shipment : %v", combinedShipment.OrdersList)
		}
	}
	c <- true
}
