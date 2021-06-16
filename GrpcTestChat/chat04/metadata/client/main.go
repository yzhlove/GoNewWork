package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-test-chat/chat04/metadata/proto"
	"io"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := proto.NewHelloClient(conn)

	data := metadata.Pairs("event", "client", "event", time.Now().Format(time.RFC3339))
	ctx := metadata.NewOutgoingContext(context.Background(), data)
	ctx = metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k2", "v3", "k3", "v3")

	var header, trailer metadata.MD

	res, err := cc.Echo(ctx, &proto.Empty{}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("res.value => ", res.Str)

	log.Println("============================== 0")
	show(header)
	log.Println()
	show(trailer)
	log.Println("============================== 0")

	st, err := cc.LoopEcho(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		for {
			if s, err := st.Recv(); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			} else {

				log.Println("==============================")
				header, _ = st.Header()
				trailer = st.Trailer()
				show(header)
				log.Println()
				show(trailer)
				log.Println("s ==> ", s.Str)
				log.Println("==============================")
			}
		}

		log.Println("============================== 4")
		header, _ = st.Header()
		trailer = st.Trailer()
		show(header)
		log.Println()
		show(trailer)
		log.Println("============================== 4")

	}()

	for _, a := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"} {
		st.Send(&proto.String{Str: a})
		time.Sleep(time.Second)
	}

	log.Println("============================== 2")
	header, _ = st.Header()
	trailer = st.Trailer()
	show(header)
	log.Println()
	show(trailer)
	log.Println("============================== 2")

	st.CloseSend()

	log.Println("============================== 3")
	header, _ = st.Header()
	trailer = st.Trailer()
	show(header)
	log.Println()
	show(trailer)
	log.Println("============================== 3")

}

func show(md metadata.MD) {
	for k, v := range md {
		for _, x := range v {
			log.Printf("k = %v v = %v", k, x)
		}
	}
}
