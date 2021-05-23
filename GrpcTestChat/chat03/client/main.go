package main

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat03/proto"
	"io"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {

	//testAddedOrder()
	//testGetOrder()

	//testSearchOrder()

	//testUpdatedOrder()

	testProcessOrder()

}

func testAddedOrder() {

	client, clos := obtainConn()
	defer clos.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	order := &proto.Order{
		Id:          "101",
		Items:       []string{"IPhone XS", "Mac Book Pro"},
		Destination: "San Jose,CA",
		Price:       2300.00,
	}

	if res, err := client.AddOrder(ctx, order); err != nil {
		log.Fatalf("added order error: %v", err)
	} else {
		log.Println("added order ID response ->", res.Value)
	}
}

func testGetOrder() {
	client, clos := obtainConn()
	defer clos.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.GetOrder(ctx, &proto.StringValue{Value: "102000"})
	if err != nil {
		log.Printf("get order error:%v", err)
	} else {
		log.Println("get order info ->", res)
	}

	res, err = client.GetOrder(ctx, &proto.StringValue{Value: "101"})

	if err != nil {
		log.Printf("get order error:%v", err)
	} else {
		log.Println("get order info ->", res)
	}

}

//testSearchOrder 服务端推流示例
func testSearchOrder() {
	client, clos := obtainConn()
	defer clos.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.SearchOrders(ctx, &proto.StringValue{Value: "Google"})
	if err != nil {
		log.Fatalf("get search order stream error: %v ", err)
	}

	for {
		result, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Print("get search order stream EOF!")
				break
			}
			log.Fatalf("search order error: %v", err)
		}

		log.Print("search result order ->", result)
	}
}

//testUpdatedOrder 客户端推流示例
func testUpdatedOrder() {
	client, clos := obtainConn()
	defer clos.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	order1 := proto.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	order2 := proto.Order{Id: "103", Items: []string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	order3 := proto.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}

	stream, err := client.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("updated order stream error: %v", err)
	}

	for _, order := range []proto.Order{order1, order2, order3} {
		if err := stream.Send(&order); err != nil {
			log.Fatalf("stream send error:%v ", err)
		}
	}

	//关闭接收方
	if res, err := stream.CloseAndRecv(); err != nil {
		if errors.Is(err, io.EOF) {
			log.Print("stream EOF!")
		}
		log.Printf("stream close and recv error: %v", err)
	} else {
		log.Print("stream close succeed:", res.Value)
	}

	//关闭发送方
	//if err := stream.CloseSend(); err != nil {
	//	log.Fatalf("stream close send error: %v", err)
	//} else {
	//	log.Print("stream send close Ok.")
	//}

}

//testProcessOrder 客户端服务端双向推流示例
func testProcessOrder() {
	client, clos := obtainConn()
	defer clos.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("process orders stream error: %v", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Print("process orders stream.recv EOF!")
					break
				}
			}

			log.Print("combined shipment :", resp.Id, resp.Status, resp.OrderList)
		}
		log.Print("stream recv routine exit...")
	}()

	for k, v := range []string{"102", "103", "104", "105"} {
		if err := stream.Send(&proto.StringValue{Value: v}); err != nil {
			log.Printf("process orders stream send index %d ID: %s error: %v ", k, v, err)
		}
	}

	if err := stream.Send(&proto.StringValue{Value: "101"}); err != nil {
		log.Printf("process orders stream send  error: %v ", err)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("stream send close error: %v", err)
	}

	log.Print("close send stream ....")
	time.Sleep(time.Second)
	log.Print("close send stream Ok.")
}

func obtainConn() (proto.OrderManagementClient, io.Closer) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	client := proto.NewOrderManagementClient(conn)

	return client, conn
}
