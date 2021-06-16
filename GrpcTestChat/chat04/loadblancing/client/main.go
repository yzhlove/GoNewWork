package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	echo "grpc-test-chat/chat04/loadblancing/proto"
	"log"
	"time"
)

var address = []string{"localhost:50051", "localhost:50052"}

func main() {

	//pick_first模式，如果第一个addr可用，则只使用第一个addr，如果第一个不可用，则使用下一个。
	pickFirstConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", "example", "example.com"),
		//grpc.WithBalancerName("pick_first"),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, grpc.PickFirstBalancerName)),
		grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer pickFirstConn.Close()

	makeRpc(pickFirstConn, 10)

	log.Println("=====================================================")

	randrobinConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", "example", "example.com"),
		//grpc.WithBalancerName("round_robin"),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer randrobinConn.Close()

	makeRpc(randrobinConn, 10)

}

func Echo(c echo.HelloClient, msg string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Echo(ctx, &echo.String{Str: msg})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("resp value -> ", resp.Str)
}

func makeRpc(cc *grpc.ClientConn, n int) {
	c := echo.NewHelloClient(cc)
	for i := 0; i < n; i++ {
		Echo(c, "this is example/load_balancing")
	}
}

type exampleResolverBuilder struct{}

func (e *exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	example := &exampleResolver{
		target:       target,
		cc:           cc,
		addressStore: map[string][]string{"example.com": address},
	}
	example.start()
	return example, nil
}

func (e *exampleResolverBuilder) Scheme() string {
	return "example"
}

type exampleResolver struct {
	target       resolver.Target
	cc           resolver.ClientConn
	addressStore map[string][]string
}

func (e *exampleResolver) start() {
	address := e.addressStore[e.target.Endpoint]
	updated := make([]resolver.Address, len(address))

	for i, s := range address {
		updated[i] = resolver.Address{Addr: s}
	}

	e.cc.UpdateState(resolver.State{Addresses: updated})
}

func (e *exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (e *exampleResolver) Close() {}

func init() {
	resolver.Register(&exampleResolverBuilder{})
}
