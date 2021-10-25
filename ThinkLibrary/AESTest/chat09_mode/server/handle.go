package main

import (
	"context"
	"think-library/AESTest/chat09_mode/conf"
	"think-library/AESTest/chat09_mode/dh"
	"think-library/AESTest/chat09_mode/pack"
	"think-library/AESTest/chat09_mode/pb"
)

type handleFunc func(ctx context.Context, msg pack.Packet) pack.Msg

var MsgCenter map[uint16]handleFunc

func auth(ctx context.Context, msg pack.Packet) pack.Msg {
	key, err := dh.GenerateKey()
	if err != nil {
		return &pb.Error{Error: err.Error()}
	}

	return nil
}

func init() {
	MsgCenter = make(map[uint16]handleFunc)
	MsgCenter[conf.SystemAuth] = auth
}
