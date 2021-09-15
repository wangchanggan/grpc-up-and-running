package main

import (
	"context"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "ordermgt/client/ecommerce"
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

	// 创建到服务器端的连接并初始化调用服务的客户端存根。
	client := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Add Order
	order1 := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	res, _ := client.AddOrder(ctx, &order1)
	if res != nil {
		log.Print("AddOrder Response -> ", res.Value)
	}

	// Get Order
	// 调用客户端存根的GetOrder方法，实现对远程方法的调用。
	// 这时会得到一个order消息作为响应，其中包含服务定义中使用protocol buffers所定义的订单信息。
	retrievedOrder, err := client.GetOrder(ctx, &wrapper.StringValue{Value: "106"})
	log.Print("GetOrder Response -> : ", retrievedOrder)

	// Search Order : Server streaming scenario
	// SearchOrders 方法返回OrderManagenent_SearchOrdersClient 的客户端流，它有一个名为Recv的方法。
	searchStream, _ := client.SearchOrders(ctx, &wrapper.StringValue{Value: "Google"})
	for {
		// 调用客户端流的Recv方法，逐个检索Order响应。
		searchOrder, err := searchStream.Recv()
		// 当发现流结束的时候，Recv会返回io.EOF。
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			log.Print("Search Result : ", searchOrder)
		}
	}

	// =========================================
	// Update Orders : Client streaming scenario
	// 客户端可以通过客户端流引用，借助updateStream.Send方法发送多条消息。
	// 一旦所有消息都以流的形式发送出去，客户端就可以将流标记为已完成，并接收来自服务器端的响应。
	// 这是通过流引用的CloseAndRecv方法实现的。
	updOrder1 := pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	updOrder2 := pb.Order{Id: "103", Items: []string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	updOrder3 := pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}

	// 调用UpdateOrders远程方法。
	updateStream, err := client.UpdateOrders(ctx)
	// 处理与UpdateOrders相关的错误。
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", client, err)
	}

	// Updating order 1
	// 通过客户端流发送订单更新的请求。
	if err := updateStream.Send(&updOrder1); err != nil {
		// 处理在发送消息到流时发生的错误。
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	}

	// Updating order 2
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	}

	// Updating order 3
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder3, err)
	}

	// 关闭流并接收响应。
	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updateRes)

	// =========================================
	// Process Order : Bi-di streaming scenario
	// 当客户端通过OrderManagement对象调用ProcessOrders方法时，它会得到一个对流的引用(streamProcOrder)，
	// 这个引用可以用来发送消息到服务器端，也能读取来自服务器端的消息。
	streamProcOrder, err := client.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("%v.ProcessOrders(_) = _, %v", client, err)
	}

	// 向服务发送消息。
	if err := streamProcOrder.Send(&wrapper.StringValue{Value: "102"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "102", err)
	}

	if err := streamProcOrder.Send(&wrapper.StringValue{Value: "103"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "103", err)
	}

	if err := streamProcOrder.Send(&wrapper.StringValue{Value: "104"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "104", err)
	}

	// 创建Goroutines所使用的通道。
	channel := make(chan struct{})
	// 使用Goroutines调用函数，以便并行读取来自服务的消息。
	go asncClientBidirectionalRPC(streamProcOrder, channel)
	// 模拟向服务发送消息的延迟。
	time.Sleep(time.Millisecond * 1000)

	if err := streamProcOrder.Send(&wrapper.StringValue{Value: "101"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "101", err)
	}

	// 为客户端流标记流的结束。
	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}
	channel <- struct{}{}

	/*客户端可以并发读取和写人同一个流， 输入流和输出流可以独立进行操作。
	这里所展示的是稍微复杂的示例，它展现了双向流RPC模式的威力。
	流的操作完全独立，客户端和服务器端可以按照任意顺序进行读取和写人，理解这一点非常重要。
	且建立连接，客户端和服务器端之间的通信模式就完全取决于客户端和服务器端本身。*/
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan struct{}) {
	for {
		// 在客户端读取服务的消息。
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		// 该条件探测流是否已经结束。
		if errProcOrder == io.EOF {
			break
		}
		log.Printf("Combined shipment : %v", combinedShipment.OrdersList)
	}
	<-c
}
