package pack

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	packSize = 4 // 包体长度
	msgSize  = 2 // 消息头长度
)

type Msg interface {
	Size() int
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func Pack(msgID uint16, msg Msg) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, binary.LittleEndian, uint32(msgSize+msg.Size())); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, msgID); err != nil {
		return nil, err
	}

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type Packet []byte

func (p Packet) Id() uint16 {
	return binary.LittleEndian.Uint16(p)
}

func (p Packet) Unpack(msg Msg) error {
	var data = make([]byte, len(p)-msgSize)
	buf := bytes.NewReader(p[msgSize:])

	if err := binary.Read(buf, binary.LittleEndian, &data); err != nil {
		return err
	}

	return msg.Unmarshal(data)
}

type Parser struct {
	Conn io.ReadWriter
	err  error
}

func NewParser(conn io.ReadWriter) *Parser {
	return &Parser{Conn: conn}
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) Next() (Packet, bool) {
	var size uint32
	if err := binary.Read(p.Conn, binary.LittleEndian, &size); err != nil {
		p.err = err
		return nil, false
	}

	data := make([]byte, size)
	if err := binary.Read(p.Conn, binary.LittleEndian, &data); err != nil {
		p.err = err
		return nil, false
	}
	p.err = nil
	return data, true
}
