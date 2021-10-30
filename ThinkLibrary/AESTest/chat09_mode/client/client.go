package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"think-library/AESTest/chat09_mode/conf"
	"think-library/AESTest/chat09_mode/dh"
	"think-library/AESTest/chat09_mode/pack"
	"think-library/AESTest/chat09_mode/pb"
	"think-library/AESTest/chat09_mode/secret"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", conf.HostName)
	if err != nil {
		toErr("DialTCP", err)
	}

	ctx := NewContext(conn)
	
	auth := &pb.Auth{Msg: "client pubKey", PubKey: ctx.Dh.PubKey.Bytes()}
	if err := ctx.Send(conf.SystemAuth, auth); err != nil {
		toErr("Send Auth", err)
	}

	ctx.ReaderMsg()

}

func toErr(prefix string, err error) {
	panic(fmt.Sprintf("[%s] panic:%v", prefix, err))
}

func toShow(msg string) {
	log.Println("[client message]", msg)
}

type Context struct {
	Conn  net.Conn
	Dh    *dh.DhKey
	Cert  secret.EncoderDecoder
	msgCh chan pack.Packet
}

func NewContext(conn net.Conn) *Context {
	ctx := &Context{Conn: conn}
	s, err := dh.GenerateKey()
	if err != nil {
		toErr("GenerateKey", err)
	}
	ctx.Dh = s

	c, err := secret.NewAes(conf.DefineKey)
	if err != nil {
		toErr("NewAes", err)
	}
	ctx.Cert = c

	ctx.msgCh = make(chan pack.Packet, 1)

	return ctx
}

func (ctx *Context) Send(msgID uint16, msg pack.Msg) error {
	if msg != nil {
		data, err := msg.Marshal()
		if err != nil {
			toErr("Marshal", err)
		}

		if data, err = ctx.Cert.Encode(data); err != nil {
			toErr("Encode", err)
		}

		if data, err = pack.Pack(msgID, data); err != nil {
			toErr("Pack", err)
		}

		if _, err = ctx.Conn.Write(data); err != nil {
			return err
		}
		return nil
	}
	return errors.New("msg not empty")
}

func (ctx *Context) handlerMsg() {
	for p := range ctx.msgCh {
		switch p.Id() {
		case conf.SystemAuth:
			authHandle(ctx, p)
		case conf.SystemNotFound:
			notFoundHandle(ctx, p)
		case conf.UserLetter:
			echoHandle(ctx, p)
		default:
			toShow(fmt.Sprintf("[%d] not found message", p.Id()))
		}
	}
}

func (ctx *Context) req(data []byte) pack.Packet {
	if len(data) > 0 {
		ret, err := ctx.Cert.Decode(data[pack.MsgSize:])
		if err != nil {
			toErr("Decode", err)
		}
		return append(data[:pack.MsgSize], ret...)
	}
	return nil
}

func (ctx *Context) ReaderMsg() {
	p := parser{conn: ctx.Conn}
	go ctx.handlerMsg()

	for {
		data, err := p.read()
		if err != nil {
			toErr("read", err)
		}
		req := ctx.req(data)
		if req != nil {
			ctx.msgCh <- req
		}
	}
}

type parser struct {
	conn net.Conn
}

func (p *parser) read() ([]byte, error) {
	var size uint32
	if err := binary.Read(p.conn, binary.LittleEndian, &size); err != nil {
		return nil, err
	}

	if size == 0 {
		return nil, errors.New("read conn size not is zero")
	}

	buf := make([]byte, size)
	//if err := binary.Read(p.conn, binary.LittleEndian, &data); err != nil {
	//	return nil, err
	//}

	if _, err := io.ReadFull(p.conn, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func correctKey(data []byte) []byte {
	x := make([]byte, 32)
	copy(x, data)
	return x
}

func authHandle(ctx *Context, reader pack.Packet) {

	tbl := &pb.Auth{}
	if err := reader.Unpack(tbl); err != nil {
		toErr("Unpack", err)
	}

	toShow(tbl.Msg)

	key, err := dh.ComputeKey(big.NewInt(0).SetBytes(tbl.PubKey), ctx.Dh.PriKey)
	if err != nil {
		toErr("ComputeKey", err)
	}

	cert, err := secret.NewAes(correctKey(key.Bytes()))
	if err != nil {
		toErr("NewAes", err)
	}

	// 更新客户端的密钥
	ctx.Cert = cert

	letter := &pb.Letter{Recipient: "Server", Sender: "Client", Content: []byte("key succeed.")}

	if err := ctx.Send(conf.UserLetter, letter); err != nil {
		toErr("Send", err)
	}
}

func notFoundHandle(ctx *Context, reader pack.Packet) {

	tbl := &pb.NotFoundResp{}
	if err := reader.Unpack(tbl); err != nil {
		toErr("Unpack", err)
	}

	toShow("NotFound:" + tbl.Msg)

}

func echoHandle(ctx *Context, reader pack.Packet) {

	tbl := &pb.Letter{}
	if err := reader.Unpack(tbl); err != nil {
		toErr("Unpack", err)
	}

	toShow(fmt.Sprintf("%s:%s [%s]", tbl.Sender, tbl.Recipient, string(tbl.Content)))

	time.Sleep(time.Second * 10)

	tbl.Sender = "Client"
	tbl.Recipient = "Server"

	if err := ctx.Send(conf.UserLetter, tbl); err != nil {
		toErr("Send", err)
	}
}
