package main

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat02/proto"
	"log"
	"net"
	"time"
)

var address = "localhost:50051"

func main() {

	stat := make(chan struct{})
	go startTest(stat)

	client()
	time.Sleep(time.Second)

	fmt.Println("Ok.")

}

func startTest(stat chan struct{}) {
	go server()
	<-stat
}

func client() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc dial error", err)
	}
	defer conn.Close()

	cli := proto.NewProductInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := cli.AddProduct(ctx, &proto.Product{Name: "Apple IPhoneX", Description: "This is apple phone", Price: 6999.00})
	if err != nil {
		log.Fatalln("add product err ", err)
	}
	log.Printf("pprodcut id => %s", result.Value)

	product, err := cli.GetProduct(ctx, &proto.ProductID{Value: result.Value})
	if err != nil {
		log.Fatalln("get product err ", err)
	}
	log.Printf("product:%v ", product.String())
}

type service struct {
	productMap map[string]*proto.Product
}

func (s *service) AddProduct(ctx context.Context, in *proto.Product) (*proto.ProductID, error) {

	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error generate product id", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*proto.Product, 1<<4)
	}
	s.productMap[in.Id] = in
	log.Printf("product %v : %v - Added.", in.Id, in.Name)
	return &proto.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *service) GetProduct(ctx context.Context, in *proto.ProductID) (*proto.Product, error) {
	product, ok := s.productMap[in.Value]
	if ok && product != nil {
		log.Printf("product %v : %v - Retrieved.", product.Id, product.Name)
		return product, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Product does not exists.", in.Value)
}

func server() {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("listener error ", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(1024))
	proto.RegisterProductInfoServer(s, &service{})

	if err := s.Serve(l); err != nil {
		log.Fatalln("server error ", err)
	}
}
