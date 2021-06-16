package main

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat04/compression/proto"
	"log"
	"time"
)

func main() {

	cc, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	cli := proto.NewHelloClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := cli.Echo(ctx, &proto.String{Str: "-1"})
	if err != nil {

		errcode := status.Code(err)

		if errcode == codes.InvalidArgument {

			log.Printf("Invalid Arguments Error: %s", errcode)

			errorSatus := status.Convert(err)

			for _, d := range errorSatus.Details() {

				switch info := d.(type) {
				case *errdetails.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s ", info)
				default:
					log.Printf("Unexcept error type: %s", info)
				}

			}

		} else {
			log.Printf("unhandler error: %s ", errcode)
		}

	}

	log.Print("resp -> ", resp.Str)

}
