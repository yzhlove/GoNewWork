package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"grpc-test-chat/chat04/protocol/common"
	"grpc-test-chat/chat04/protocol/proto"
	"io"
	"log"
	"time"
)

func main() {

	//testInterceptAddedOrder()
	//testInterceptGetOrder()
	//testInterceptSearchOrder()
	//testInterceptUpdatedOrder()
	testInterceptProcessOrder()

}

func testInterceptAddedOrder() {

	client, conn := obtainedConn()
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

func testInterceptGetOrder() {
	client, conn := obtainedConn()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.GetOrder(ctx, &common.String{Value: "101"})
	if err != nil {
		log.Printf("get order err:%v ", err)
		return
	}
	log.Printf("get order info: %v ", res)
}

func testInterceptSearchOrder() {
	client, conn := obtainedConn()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.SearchOrders(ctx, &common.String{Value: "Google"})
	if err != nil {
		log.Printf("search order error: %v ", err)
		return
	}

	for {
		if order, err := stream.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				log.Print("search order stream eof")
				break
			}
			log.Fatalf("search order stream error:%v ", err)
		} else {
			log.Printf("search order -> %v ", order)
		}
	}
}

func testInterceptUpdatedOrder() {
	client, conn := obtainedConn()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	order1 := proto.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	order2 := proto.Order{Id: "103", Items: []string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	order3 := proto.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}

	stream, err := client.UpdatedOrders(ctx)
	if err != nil {
		log.Fatalf("stream error:%v", err)
	}

	for _, order := range []proto.Order{order2, order1, order3} {
		if err := stream.Send(&order); err != nil {
			log.Fatalf("send updated order error: %v", err)
		}
	}

	if rt, err := stream.CloseAndRecv(); err != nil {
		if errors.Is(err, io.EOF) {
			log.Printf("close recv eof")
			return
		}
		log.Printf("close recv error: %v ", err)
	} else {
		log.Printf("close recv succeed: %v ", rt.Value)
	}

}

func testInterceptProcessOrder() {

	client, conn := obtainedConn()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stream, err := client.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("proccess stream error: %v", err)
	}

	go func() {
		dept, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("process order eof")
				return
			}
			log.Printf("process order stream recv error: %v ", err)
		}
		log.Printf("stream recv info:%v ", dept)
	}()

	for k, v := range []string{"102", "102", "103", "101", "105", "106", "102", "103", "101", "105", "106"} {
		if err := stream.Send(&common.String{Value: v}); err != nil {
			log.Printf("send error: k: %v err: %v", k, err)
		}
	}

	time.Sleep(time.Second)

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("stream send close error:%v ", err)
	}

	time.Sleep(time.Second)
	log.Print("Ok.")
}

func obtainedConn() (proto.OrderManagementClient, io.Closer) {
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
