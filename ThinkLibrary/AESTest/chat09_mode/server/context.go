package main

import (
	"context"
	"errors"
	"think-library/AESTest/chat09_mode/conf"
	"think-library/AESTest/chat09_mode/dh"
	"think-library/AESTest/chat09_mode/pack"
	"think-library/AESTest/chat09_mode/secret"
)

type Context struct {
	Crt   secret.EncoderDecoder
	Dh    *dh.DhKey
	Fail  bool
	msgCh chan []byte
	ctx   context.Context
}

func NewContext() *Context {
	c := &Context{}

	s, err := secret.NewAes(conf.DefineKey)
	if err != nil {
		toErr("NewAes", err)
	}
	c.Crt = s

	d, err := dh.GenerateKey()
	if err != nil {
		toErr("GenerateKey", err)
	}
	c.Dh = d

	return c
}

func (c *Context) req(data []byte) pack.Packet {
	if len(data) > 0 {
		ret, err := c.Crt.Decode(data)
		if err != nil {
			toErr("Decode", err)
		}
		return ret
	}
	return nil
}

func (c *Context) send(data []byte) error {
	if len(data) > 0 {
		select {
		case c.msgCh <- data:
		default:
			return errors.New("send message error")
		}
	}
	return nil
}

func (c *Context) failed(err error) []byte {
	c.Fail = true
	//data, err := pack.Pack(conf.SystemError, &pb.ErrorResp{Error: err.Error()})
	//if err != nil {
	//	toErr("pack", err)
	//}

	//data, err := c.Crt.Encode(data)

	return nil
}

func (c *Context) succeed(msgID uint16, msg pack.Msg) []byte {
	if msg != nil {
		c.Fail = false

	}
	return nil
}
