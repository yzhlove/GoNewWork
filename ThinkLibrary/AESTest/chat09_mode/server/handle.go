package main

import (
	"think-library/AESTest/chat09_mode/conf"
	"think-library/AESTest/chat09_mode/pack"
	"think-library/AESTest/chat09_mode/pb"
)

type handFunc func(ctx *Context, reader pack.Packet) []byte

func auth(ctx *Context, reader pack.Packet) []byte {

	tbl := &pb.Auth{}
	if err := reader.Unpack(tbl); err != nil {
		//return ctx.failed()
	}

	return nil
}

func notFound(ctx *Context, reader pack.Packet) []byte {

	return nil
}

var api map[uint16]handFunc

func init() {
	api[conf.SystemAuth] = auth
}

func Get(msgID uint16) handFunc {
	if ret, ok := api[msgID]; ok {
		return ret
	}
	return notFound
}
