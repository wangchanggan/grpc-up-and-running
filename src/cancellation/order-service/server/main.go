package main

import (
	pb "cancellation/order-service/order-service-gen"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
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

// Bi-directional Streaming RPC
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {

	batchMarker := 1

	// Cancel from Server Side
	//_, cancel := context.WithCancel(stream.Context())
	//cancel()

	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		// You can determine whether the current RPC is cancelled by the other party.
		// 当某一方取消RPC之后，另一方可以通过检查context来确定这一点
		// 检查当前的上下文是否已经取消。
		if stream.Context().Err() == context.Canceled {
			log.Printf(" Context Cacelled for this stream: -> %s", stream.Context().Err())
			log.Printf("Stopped processing any more order of this stream!")
			return stream.Context().Err()
		}
		orderId, err := stream.Recv()
		log.Println("Reading Proc order ... ", orderId)

		if err == io.EOF {
			// Client has sent all the messages
			// Send remaining shipments
			log.Println("EOF for ", orderId)

			for _, comb := range combinedShipmentMap {
				stream.Send(&comb)
			}
			return nil
		}
		destination := orderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := orderMap[orderId.GetValue()]
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (orderMap[orderId.GetValue()].Destination), Status: "Processed!",}
			ord := orderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Print("Shipping : ", comb.Id, " -> ", len(comb.OrdersList))
				stream.Send(&comb)
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
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
