package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"think-library/AESTest/chat09_mode/conf"
	"think-library/AESTest/chat09_mode/dh"
	"think-library/AESTest/chat09_mode/pack"
	"think-library/AESTest/chat09_mode/pb"
	"think-library/AESTest/chat09_mode/secret"
)

type handFunc func(ctx *Context, reader pack.Packet) []byte

func auth(ctx *Context, reader pack.Packet) []byte {

	tbl := &pb.Auth{}
	if err := reader.Unpack(tbl); err != nil {
		return ctx.failed(fmt.Errorf("auth Unpack error:%w", err))
	}

	log.Println("auth->", tbl.Msg)

	newKey, err := dh.ComputeKey(ctx.Dh.PriKey, big.NewInt(0).SetBytes(tbl.PubKey))
	if err != nil {
		return ctx.failed(fmt.Errorf("auth ComputeKey error:%w", err))
	}

	ret := &pb.Auth{Msg: "Server PubKey", PubKey: ctx.Dh.PubKey.Bytes()}
	// 回复客户端的消息仍旧采用旧的密钥加密消息
	resp := ctx.succeed(conf.SystemAuth, ret)

	// 服务端采用新的密钥
	ctx.Crt, err = secret.NewAes(newKey.Bytes())
	if err != nil {
		toErr("newKey error", err)
	}

	return resp
}

func notFound(ctx *Context, reader pack.Packet) []byte {

	resp := &pb.NotFoundResp{Msg: fmt.Sprintf("not found message:%d", reader.Id())}

	return ctx.succeed(conf.SystemNotFound, resp)
}

func echo(ctx *Context, reader pack.Packet) []byte {

	tbl := &pb.Letter{}
	if err := reader.Unpack(tbl); err != nil {
		ctx.failed(fmt.Errorf("echo error:%w ", err))
	}

	resp := &pb.Letter{}
	resp.Recipient = "Client"
	resp.Sender = "Server"
	resp.Content = []byte(strings.ToUpper(string(tbl.Content)))

	return ctx.succeed(conf.UserLetter, resp)
}

var api map[uint16]handFunc

func init() {
	api = make(map[uint16]handFunc)
	api[conf.SystemAuth] = auth
	api[conf.UserLetter] = echo
}

func Get(msgID uint16) handFunc {
	if ret, ok := api[msgID]; ok {
		return ret
	}
	return notFound
}
