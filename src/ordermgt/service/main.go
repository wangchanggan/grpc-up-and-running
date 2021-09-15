package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	/*"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"*/
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"io"
	"log"
	"net"
	pb "ordermgt/service/ecommerce"
	"strings"
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
func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)
	orderMap[orderReq.Id] = *orderReq
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Simple RPC
// 作为GetOrder方法的输入，单个订单ID (String)用来组成请求，服务器端找到订单并以order消息(order结构体)的形式进行响应。
// order 消息可以和nil错误一起返回，从而告诉gRPC，已经处理完RPC, 可以将Order返回到客户端了。
func (s *server) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord, exists := orderMap[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}

// Server-side Streaming RPC
// SearchOrders 方法有两个参数，分别是字符串类型的searchQuery和用来写人响应的特殊参数OrderManagement_SearchOrdersServer。
// OrderManagement_SearchOrdersServer 是流的引用对象，可以写入多个响应。
// 业务逻辑是找到匹配的订单，并通过流将其依次发送出去。当找到新的订单时，使用流引用对象的Send(...)方法将其写入流。
// 一旦所有响应都写到了流中，就可以通过返回nil来标记流已经结束，服务器端的状态和其他trailer元数据会发送给客户端。
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {

	for key, order := range orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			// 查找匹配的订单。
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				// 通过流发送匹配的订单。
				err := stream.Send(&order)
				if err != nil {
					// 检查在将消息以流的形式发送给客户端的过程中可能出现的错误。
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

// Client-side Streaming RPC
// UpdateOrders 方法有一个orderManagenent_UpdateOrdersServer参数，它是客户端传入消息流的引用对象。
// 通过调用该对象的Recv方法来读取消息。根据业务逻辑，可以读取其中一些消息，也可以读取所有的消息。
// 服务只需调用OrderManagenent_UpdateOrdersServer对象的SendAndClose方法就可以发送响应，它同时也标记服务器端消息终结了流。
// 如果要提前停止读取客户端流，那么服务器端应该取消客户端流，这样客户端就知道停止生产消息了
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs : "
	for {
		// 从客户端流中读取消息。
		order, err := stream.Recv()
		// 检查流是否已经结束。
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}

		if err != nil {
			return err
		}
		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		ordersStr += order.Id + ", "
	}
}

// Bi-directional Streaming RPC
// ProcessOrders 方法有一个OrderManagement_ProcessOrdersServer参数，它是客户端和服务器端之间消息流的对象引用。
// 借助这个流对象，服务器端可以读取客户端以流的方式发送的消息，也能写入服务器端的流消息并返回给客户端。
// 传入的消息流可以通过该引用对象的Recv方法来读取。
// 在ProcessOrders 方法中，服务可在持续读取传入消息流的同时，使用Send方法将消息写入同一个流中。
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {

	batchMarker := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		// 从传入的流中读取订单ID。
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order : %s", orderId)
		// 持续读取，直到流结束为止。
		if err == io.EOF {
			// Client has sent all the messages
			// Send remaining shipments
			log.Printf("EOF : %s", orderId)
			for _, shipment := range combinedShipmentMap {
				// 当流结束时，将所有剩余的发货组合发送给客户端。
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			// 通过返回nil标记服务器端流已经结束。
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
			comShip := pb.CombinedShipment{Id: "cmb - " + (orderMap[orderId.GetValue()].Destination), Status: "Processed!"}
			ord := orderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		// 按批次处理订单。当达到该批次的规模时，将所有已创建的发货组合以流的形式发送给客户端。
		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrdersList))
				// 将发货组合写人流中。
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}

	/*这里是基于订单ID来处理传人的订单的，当创建新的发货组合后，服务会将其写人相同的流中。
	这与客户端流RPC模式不同，当时服务通过SendAndClose方法写人流并将其关开。
	当发现客户端流已经结束时， 发送nil标记服务器端流的结束。*/
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
	// reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}
