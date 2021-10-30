package main

import (
	"context"
	"errors"
	"think-library/AESTest/chat09_mode/conf"
	"think-library/AESTest/chat09_mode/dh"
	"think-library/AESTest/chat09_mode/pack"
	"think-library/AESTest/chat09_mode/pb"
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
	c.msgCh = make(chan []byte, 1)
	return c
}

func (c *Context) req(data []byte) pack.Packet {
	if len(data) > 0 {
		ret, err := c.Crt.Decode(data[pack.MsgSize:])
		if err != nil {
			toErr("Decode", err)
		}
		return append(data[:pack.MsgSize], ret...)
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

	tbl := &pb.ErrorResp{Error: err.Error()}

	data, err := tbl.Marshal()
	if err != nil {
		toErr("Marshal", err)
	}

	if data, err = c.Crt.Encode(data); err != nil {
		toErr("Encode", err)
	}

	if data, err = pack.Pack(conf.SystemError, data); err != nil {
		toErr("Pack", err)
	}

	return data
}

func (c *Context) succeed(msgID uint16, msg pack.Msg) []byte {
	if msg != nil {
		c.Fail = false

		data, err := msg.Marshal()
		if err != nil {
			toErr("msg.Marshal", err)
		}

		if data, err = c.Crt.Encode(data); err != nil {
			toErr("Encode", err)
		}

		if data, err = pack.Pack(msgID, data); err != nil {
			toErr("Pack", err)
		}

		return data
	}
	return nil
}
