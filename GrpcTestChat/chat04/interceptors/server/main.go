package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat04/protocol/common"
	"grpc-test-chat/chat04/protocol/proto"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const port = ":50051"

var batchSize = 2

func main() {

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryServerIntercept),
		grpc.StreamInterceptor(ServerStreamIntercept))

	sev := &server{orderMap: make(map[string]proto.Order, 1<<4)}

	proto.RegisterOrderManagementServer(service, sev)
	if err := service.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	orderMap map[string]proto.Order
}

/*
	AddOrder(context.Context, *Order) (*common.String, error)
	GetOrder(context.Context, *common.String) (*Order, error)
	SearchOrders(*common.String, OrderManagement_SearchOrdersServer) error
	UpdatedOrders(OrderManagement_UpdatedOrdersServer) error
	ProcessOrders(OrderManagement_ProcessOrdersServer) error
*/

func (s *server) AddOrder(ctx context.Context, orderReq *proto.Order) (*common.String, error) {
	s.orderMap[orderReq.Id] = *orderReq
	log.Println("Added Order ID: ", orderReq.Id)
	return &common.String{Value: "Order Added ID:" + orderReq.Id}, nil
}

func (s *server) GetOrder(ctx context.Context, orderId *common.String) (*proto.Order, error) {
	if order, ok := s.orderMap[orderId.Value]; ok {
		return &order, nil
	}
	return nil, status.New(codes.NotFound, "Not Found Order ID:"+orderId.Value).Err()
}

func (s *server) SearchOrders(tag *common.String, stream proto.OrderManagement_SearchOrdersServer) error {
	orders := make([]*proto.Order, 0, len(s.orderMap))
	for key, order := range s.orderMap {
		for _, item := range order.Items {
			if strings.Contains(item, tag.Value) {
				log.Print("Matching Order Found : ", key)
				orders = append(orders, &order)
				break
			}
		}
	}
	for _, order := range orders {
		if err := stream.Send(order); err != nil {
			log.Printf("search sender order id: %s error: %v ", order.Id, err)
			break
		}
	}
	return nil
}

func (s *server) UpdatedOrders(stream proto.OrderManagement_UpdatedOrdersServer) error {

	str := strings.Builder{}

	for {
		order, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return stream.SendAndClose(&common.String{Value: "Orders processed " + str.String()})
			}
			log.Printf("Updated Order Error: %v", err)
			break
		}

		s.orderMap[order.Id] = *order
		log.Printf("Updated Order Id: %v Order Message: %v", order.Id, order)

		str.WriteString(order.Id + ",")
	}

	return nil
}

func (s *server) ProcessOrders(stream proto.OrderManagement_ProcessOrdersServer) error {

	var dept = make(map[string]*proto.CombinedShipment, batchSize)
	for {
		orderId, err := stream.Recv()
		if err != nil {
			log.Printf("Reading Proc OrderMap:%s ", orderId)
			if errors.Is(err, io.EOF) {
				for _, dp := range dept {
					if err := stream.Send(dp); err != nil {
						return err
					}
				}
				return nil
			}
			log.Printf("recv error:%v ", err)
			return err
		}

		order, found := s.orderMap[orderId.Value]
		destination := "unknown-dest"
		if found {
			if order.Destination != "" {
				destination = order.Destination
			}
		}

		if _, found := dept[destination]; !found {
			dept[destination] = &proto.CombinedShipment{Id: destination, Status: "Ok."}
		}
		dept[destination].OrderList = append(dept[destination].OrderList, &order)

		if dp := dept[destination]; len(dp.OrderList) >= batchSize {
			log.Printf("send dept value:%v ", dp)
			if err := stream.Send(dp); err != nil {
				log.Printf("send comb dept data error: %v ", err)
				return err
			}
			delete(dept, destination)
		}

	}
}

//UnaryServerIntercept 服务端一元拦截器
func UnaryServerIntercept(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Print("====== [Server Intercept] ", info.FullMethod, " ", info.Server)
	log.Printf("====== [Server Intercept] message:%s ", req)

	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error:%v ", err)
	}
	log.Printf("Post RPC Message:%v ", m)
	return m, err
}

//ServerStreamIntercept 服务端流式拦截器
func ServerStreamIntercept(
	srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Print("====== [Server Stream Intercept] ", srv, " ", info.FullMethod)

	err := handler(srv, NewServerStream(ss))
	if err != nil {
		log.Printf("RPC failed with error:%v ", err)
	}
	return err
}

type stream struct {
	grpc.ServerStream
}

func (s *stream) RecvMsg(m interface{}) error {
	log.Printf("====== [Server Stream Intercept Wrapper] Recvive a messgae (Type:%T) value:%v at %v",
		m, m, time.Now().Format(time.RFC3339))
	return s.ServerStream.RecvMsg(m)
}

func (s *stream) SendMsg(m interface{}) error {
	log.Printf("====== [Server Stream Intercept Wrapper] Sender a message (Type:%T) value:%v at %v ",
		m, m, time.Now().Format(time.RFC3339))
	return s.ServerStream.SendMsg(m)
}

func NewServerStream(s grpc.ServerStream) grpc.ServerStream {
	return &stream{ServerStream: s}
}
