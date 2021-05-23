package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat03/proto"
	"io"
	"log"
	"net"
	"strings"
)

const (
	port           = ":50051"
	orderBatchSize = 3
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	sev := &server{orderMap: sourceData()}

	s := grpc.NewServer(grpc.MaxConcurrentStreams(1024))
	proto.RegisterOrderManagementServer(s, sev)

	log.Println("listen to ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
}

type server struct {
	orderMap map[string]proto.Order
}

func (s *server) AddOrder(ctx context.Context, order *proto.Order) (*proto.StringValue, error) {
	log.Printf("Order Added. ID:%s", order.Id)
	s.orderMap[order.Id] = *order
	return &proto.StringValue{Value: "Order Added:" + order.Id}, nil
}

func (s *server) GetOrder(ctx context.Context, value *proto.StringValue) (*proto.Order, error) {
	if order, ok := s.orderMap[value.Value]; ok {
		return &order, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Order does not exists. : %v", value.Value)
}

func (s *server) SearchOrders(value *proto.StringValue, stream proto.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, str := range order.Items {
			log.Print(str)
			if strings.Contains(str, value.Value) {
				if err := stream.Send(&order); err != nil {
					return fmt.Errorf("error sending message to stream:%v", err)
				}
				log.Print("Matching Order Found:" + key)
				break
			}
		}
	}
	return nil
}

func (s *server) UpdateOrders(stream proto.OrderManagement_UpdateOrdersServer) error {
	sb := strings.Builder{}
	sb.WriteString("Updated Order Ids :")
	for {
		order, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&proto.StringValue{Value: "Order processed closed" + sb.String()})
		}
		if err != nil {
			return err
		}
		s.orderMap[order.Id] = *order

		log.Printf("Order ID: %s - %s ", order.Id, "Updated")
		sb.WriteString(order.Id + ",")
	}
}

func (s *server) ProcessOrders(stream proto.OrderManagement_ProcessOrdersServer) error {
	batchMaker := 1
	var combinedShipmentMap = make(map[string]proto.CombinedShipment, 1<<4)
	for {
		orderId, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Printf("Reading Proc orderMap : %s ", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}

		if err != nil {
			log.Println(err)
			return err
		}

		destination := s.orderMap[orderId.Value].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := s.orderMap[orderId.Value]
			shipment.OrderList = append(shipment.OrderList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := proto.CombinedShipment{Id: "cmb - " + s.orderMap[orderId.Value].Destination, Status: "Processed!"}
			ord := s.orderMap[orderId.Value]
			comShip.OrderList = append(comShip.OrderList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrderList), comShip.Id)
		}

		if batchMaker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping: %s -> %d ", comb.Id, len(comb.OrderList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMaker = 0
			combinedShipmentMap = make(map[string]proto.CombinedShipment)
		} else {
			batchMaker++
		}

	}
}

func sourceData() map[string]proto.Order {
	orderMap := make(map[string]proto.Order, 5)
	orderMap["102"] = proto.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = proto.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = proto.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = proto.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = proto.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
	return orderMap
}
