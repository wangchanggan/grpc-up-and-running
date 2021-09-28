package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/resolver"
)

const (
	exampleScheme      = "example"
	exampleServiceName = "lb.example.grpc.io"
)

var addrs = []string{"localhost:50051", "localhost:50052"}

func callUnaryEcho(c ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Message)
}

func makeRPCs(cc *grpc.ClientConn, n int) {
	hwc := ecpb.NewEchoClient(cc)
	for i := 0; i < n; i++ {
		callUnaryEcho(hwc, "this is examples/load_balancing")
	}
}

func main() {
	pickfirstConn, err := grpc.Dial(
		// 使用模式和服务名创建gRPC连接。模式是通过模式解析器解析的，它是客户端应用程序的一部分。
		// 指定负载均衡算法，该算法会使用服务器端点列表中的第一个服务器端。
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName), // "example:///lb.example.grpc.io"
		grpc.WithBalancerName("pick_first"),                        // "pick_first" is the default, so this DialOption is not necessary.
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer pickfirstConn.Close()

	log.Println("==== Calling helloworld.Greeter/SayHello with pick_first ====")
	makeRPCs(pickfirstConn, 10)

	// Make another ClientConn with round_robin policy.
	roundrobinConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName), // // "example:///lb.example.grpc.io"
		// 使用轮询调度算法。
		grpc.WithBalancerName("round_robin"), // This sets the initial balancing policy.
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer roundrobinConn.Close()

	log.Println("==== Calling helloworld.Greeter/SayHello with round_robin ====")
	makeRPCs(roundrobinConn, 10)

	/*gRPC有两个默认支持的负载均衡算法: pick_first和round_robin。
	pick_first 会尝试连接第一个地址，如果能够连接成功，就会将该地址用于所有的RPC；如果失败，则会尝试下一个地址。
	round_robin 会连接所有地址，并会按顺序每次向后端发送一个RPC。*/
}

// Name resolver implementation
// 命名解析器构建器。
type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// 创建解析lb.example.grpc.io的示例解析器。
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			// 将lb.example.grpc.io解析为localhost:50051和localhost:50052。
			exampleServiceName: addrs, // "lb.example.grpc.io": "localhost:50051", "localhost:50052"
		},
	}
	r.start()
	return r, nil
}

// 为example模式创建的解析器。
func (*exampleResolverBuilder) Scheme() string { return exampleScheme } // "example"

// 命名解析器的结构。
type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}

func init() {
	resolver.Register(&exampleResolverBuilder{})
}
