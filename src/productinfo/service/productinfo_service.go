package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	// 导入刚刚通过protobuf编译器所生成的代码所在的包
	pb "productinfo/service/ecommerce"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
// server结构体是对服务器的抽象。可以通过它将服务方法附加到服务器上。
type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
// AddProduct方法以Product作为参数并返回一个ProductID。Product和ProductID结构体定义在product_info.pb.go 文件中，该文件是通过product_info.proto定义自动生成的。
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	log.Printf("Product %v : %v - Added.", in.Id, in.Name)
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
// GetProduct 方法以ProductID作为参数并返回product
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	product, exists := s.productMap[in.Value]
	if exists && product != nil {
		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
		return product, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

/*这两个方法都有一个 context参数。Context 对象包含些元数据，比如终端用户授权令牌的标识和请求的截止时间。这些元数据会在请求的生命周期内一直存在。
这两个方法都会返回一个错误以及远程方法的返回值(方法有多种返回类型)。这些错误会传播给消费者，用来进行消费者端的错误处理。*/

func main() {
	// 由gRPC服务器所绑定的TCP监听器在给定的端口(50051)上创建。
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 通过调用gRPC Go API创建新的gRPC服务器实例。
	s := grpc.NewServer()
	// 通过调用生成的API, 将之前生成的服务注册到新创建的gRPC服务器上。
	pb.RegisterProductInfoServer(s, &server{})
	// 在指定端口（50051）上开始监听传入的消息
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}