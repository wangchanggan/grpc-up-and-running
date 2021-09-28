package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "interceptors/order-service/order-service-gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
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
	log.Println("Order : ",  orderReq.Id, " -> Added")
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Simple RPC
func (s *server) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord := orderMap[orderId.Value]
	return &ord, nil
}

// Server-side Streaming RPC
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {

	for key, order := range orderMap {
		for _, itemStr := range order.Items {
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				log.Print("Matching Order Found : " + key, " -> Writing Order to the stream ... ")
				stream.Send(&order)
				break
			}
		}
	}
	return nil
}

// Client-side Streaming RPC
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}
		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}

// Bi-directional Streaming RPC
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {

	batchMarker := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Println("Reading Proc order ... ", orderId)
		if err == io.EOF {
			// Client has sent all the messages
			// Send remaining shipments

			log.Println("EOF ", orderId)

			for _, comb := range combinedShipmentMap {
				stream.Send(&comb)
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		destination := orderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := orderMap[orderId.GetValue()]
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (orderMap[orderId.GetValue()].Destination), Status: "Processed!", }
			ord := orderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Print("Shipping : " , comb.Id, " -> ", len(comb.OrdersList))
				stream.Send(&comb)
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}


// Server :: Unary Interceptor
// 服务器端一元拦截器
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Pre-processing logic
	// Gets info about the current RPC call by examining the args passed in
	// 前置处理阶段:可以在调用对应的RPC之前拦截消息。
	log.Println("======= [Server Interceptor] ", info.FullMethod)
	log.Printf(" Pre Proc Message : %s", req)


	// Invoking the handler to complete the normal execution of a unary RPC.
	// 通过UnaryHandler调用RPC方法。
	m, err := handler(ctx, req)

	// Post processing logic
	// 后置处理阶段:可以在这里处理RPC响应。
	log.Printf(" Post Proc Message : %s", m)

	// 将RPC响应发送回去。
	return m, err
}


// wrappedStream wraps around the embedded grpc.ServerStream, and intercepts the RecvMsg and
// SendMsg method call.
// 服务器端流拦截器
// wrappedStream包装嵌入的grpc.ServerStream
// 并拦截对RecvMsg函数和SendMsg的数的调用
type wrappedStream struct {   // grpc.ServerStream 的包装器流。
	grpc.ServerStream
}

// 实现包装器的RecvMsg函数，来处理流RPC所接收到的消息。
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

// 实现包装器的SendMsg函数，来处理流RPC所发送的消息。
func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

// 创建新包装器流的实例。
func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// 流拦截器的实现。
func orderServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Pre-processing
	// 前置处理阶段。
	log.Println("====== [Server Stream Interceptor] ", info.FullMethod)

	// Invoking the StreamHandler to complete the execution of RPC invocation
	// 使用包装器流调用流RPC。
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor),   // 使用gRPC服务器端注册元拦截器。
		grpc.StreamInterceptor(orderServerStreamInterceptor)) // 注册拦截器。
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
