package main

import (
	"context"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	pb "interceptors/order-service/order-service-gen"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// Setting up a connection to the server.
	// 通过传入一元拦截器作为grpc.Dial的选项，建立到服务器端的连接。
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(orderUnaryClientInterceptor), // 注册拦截器函数通过使用grpc.WithUnaryInterceptor，来在grpc.Dial操作中实现。
		grpc.WithStreamInterceptor(clientStreamInterceptor)) // 注册流拦截器。
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Add Order
	order1 := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	res, _ := c.AddOrder(ctx, &order1)
	log.Print("AddOrder Response -> ", res.Value)

	// Get Order
	retrievedOrder, err := c.GetOrder(ctx, &wrapper.StringValue{Value: "106"})
	log.Print("GetOrder Response -> : ", retrievedOrder)

	// Search Order
	searchStream, _ := c.SearchOrders(ctx, &wrapper.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			log.Print("Search Result : ", searchOrder)
		}
	}

	// Update Orders
	updOrder1 := pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	updOrder2 := pb.Order{Id: "103", Items: []string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	updOrder3 := pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}

	updateStream, _ := c.UpdateOrders(ctx)
	_ = updateStream.Send(&updOrder1)
	_ = updateStream.Send(&updOrder2)
	_ = updateStream.Send(&updOrder3)

	updateRes, _ := updateStream.CloseAndRecv()
	log.Printf("Update Orders Res : %s", updateRes)

	// Process Order
	streamProcOrder, _ := c.ProcessOrders(ctx)
	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "102"})
	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "103"})
	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "104"})

	channel := make(chan bool, 1)
	go asncClientBidirectionalRPC(streamProcOrder, channel)
	time.Sleep(time.Millisecond * 1000)

	_ = streamProcOrder.Send(&wrapper.StringValue{Value: "101"})
	_ = streamProcOrder.CloseSend()

	<-channel
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan bool) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		}
		log.Printf("Combined shipment : %v", combinedShipment.OrdersList)
	}
	c <- true
}

func orderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Pre-processor phase
	// 前置处理阶段能够在RPC请求发送至服务器端之前访问它。
	log.Println("Method : " + method)

	// Invoking the remote method
	// 通过UnaryInvoker调用RPC方法。
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Post-processor phase
	// 后置处理阶段，可以在这里处理响应结果或错误结果。
	log.Println(reply)

	// 向gRPC客户端应用程序返回错误，同时包含作为参数传递进来的答复。
	return err
}

func clientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// 前置处理阶段能够在将RPC请求发送至服务器端之前访问它。
	log.Println("======= [Client Interceptor] ", method)

	// 调用传入的streamer来获取ClientStream。
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	// 包装ClientStream，使用拦截逻辑重载其方法并返回给客户端应用程序。
	return newWrappedStream(s), nil
}

// grpc.ClientStream的包装器流。
type wrappedStream struct {
	grpc.ClientStream
}

// 拦截流RPC所接收消息的函数。
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("====== [Client Stream Interceptor] Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

// 拦截流RPC所发送消息的面数。
func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("====== [Client Stream Interceptor] Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

// 流操作拦截是通过流的包装器实现完成的，该实现中必须实现包装grpc.ClientStream的新结构。
// 这里实现了两个包装流的函数，即RecvMsg函教和SendMsg函数，分别用来拦截客户端接收及发送的流消息。
// 拦截器的注册和一元拦截器是一样的，都是通过grpc.Dial操作完成的。
func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}
