package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var writeErr = errors.New("write binary data not identical")

const (
	SystemNotDefine = 0
	SystemConnect   = 1
	SystemNonce     = 2
	UsersConnect    = 3
	UsersNonce      = 4
)

type Msg struct {
	conn io.ReadWriter
}

func (m *Msg) Write(id uint16, data []byte) error {
	dataLen := 2 + len(data)
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.LittleEndian, uint32(dataLen)); err != nil {
		return err
	}
	if err := binary.Write(buffer, binary.LittleEndian, id); err != nil {
		return err
	}
	if err := binary.Write(buffer, binary.LittleEndian, data); err != nil {
		return err
	}
	_, err := m.conn.Write(buffer.Bytes())
	return err
}

func (m *Msg) Read() (id uint16, data []byte, err error) {

	length := make([]byte, 4)
	if _, err = io.ReadFull(m.conn, length); err != nil {
		return
	}
	dataLen := binary.LittleEndian.Uint32(length)
	buf := make([]byte, dataLen)
	if _, err = io.ReadFull(m.conn, buf); err != nil {
		return
	}
	id = binary.LittleEndian.Uint16(buf[0:2])
	data = buf[2:]
	return
}
