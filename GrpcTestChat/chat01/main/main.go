package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat01/proto"
	"hash/crc32"
	"net"
	"strconv"
	"time"
)

var target = ":1234"

func main() {

	go start2()

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewUserClient(conn)
	resp, err := c.Get(context.Background(), &proto.UserContext_Req{Id: 1234567})
	if err != nil {
		panic(err)
	}
	fmt.Println("resp->", resp)

	time.Sleep(time.Second)

}

type server struct{}

func (s *server) Get(ctx context.Context, in *proto.UserContext_Req) (*proto.UserContext_Resp, error) {
	fmt.Println("in->", in.Id)
	number := crc32.ChecksumIEEE([]byte(strconv.FormatInt(int64(in.Id), 10)))
	return &proto.UserContext_Resp{Res: strconv.FormatUint(uint64(number), 10)}, nil
}

func start() error {
	l, err := net.Listen("tcp", target)
	if err != nil {
		panic(err)
	}

	ss := grpc.NewServer()
	proto.RegisterUserServer(ss, &server{})
	return ss.Serve(l)
}

func start2() error {
	l, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}
	for {
		if conn, err := l.Accept(); err != nil {
			fmt.Println("accept error:", err)
		} else {
			go func(c net.Conn) {
				reader := bufio.NewReader(conn)
				buf := make([]byte, 1024)
				for {
					n, err := reader.Read(buf)
					if err != nil {
						if n > 0 {
							fmt.Println("↓ ", string(buf[:n]))
						}
						return
					}
					fmt.Println("↓ ", string(buf[:n]))
				}
			}(conn)
		}
	}
}
