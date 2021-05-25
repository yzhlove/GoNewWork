package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat04/protocol/proto"
	"io"
	"log"
	"time"
)

func main() {

	testInterceptAddedOrder()

}

func testInterceptAddedOrder() {

	client, conn := obtainConn()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	order := &proto.Order{
		Id:          "101",
		Items:       []string{"IPhone XS", "Mac Book Pro"},
		Destination: "Made In China .",
		Price:       6999.00,
	}

	if res, err := client.AddOrder(ctx, order); err != nil {
		log.Fatalf("added order error: %v", err)
	} else {
		log.Print("added order at id:", res.Value)
	}

}



func obtainConn() (proto.OrderManagementClient, io.Closer) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(UnaryClientIntercept),
		grpc.WithStreamInterceptor(StreamClientIntercept))

	if err != nil {
		log.Fatalf("did not connect:%v ", err)
	}

	client := proto.NewOrderManagementClient(conn)
	return client, conn
}

//UnaryClientIntercept 客户端一元拦截器
func UnaryClientIntercept(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	log.Print("method:", method, " >>req:", req)
	//invoking remote func method
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Print("method:", method, " >>resp:", reply)
	return err
}

//StreamClientIntercept 流式客户端拦截器
func StreamClientIntercept(
	ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	log.Println("======== [Client Interceptor],", method, " desc:", desc)

	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		log.Fatalf("stream created error:%v ", err)
		return nil, err
	}

	return NewClientStream(s), nil
}

type stream struct {
	grpc.ClientStream
}

func (s *stream) RecvMsg(m interface{}) error {
	log.Printf("======== [Client Stream Recv], recv message type:%T value:%v at:%v ", m, m, time.Now().Format(time.RFC3339))
	return s.ClientStream.RecvMsg(m)
}

func (s *stream) SendMsg(m interface{}) error {
	log.Printf("======== [Client Stream Send], send message type:%T value:%v at:%v ", m, m, time.Now().Format(time.RFC3339))
	return s.ClientStream.SendMsg(m)
}

func NewClientStream(s grpc.ClientStream) grpc.ClientStream {
	return &stream{ClientStream: s}
}
