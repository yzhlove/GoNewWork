package pack

import "encoding/binary"

type MessageInterface interface {
	GetID() interface{}
	GetByteOrder() binary.ByteOrder
	GetData() []byte
}

type Pack struct {
	ID   uint16
	Data []byte
}

func (p *Pack) GetID() interface{} {
	return p.ID
}

func (p *Pack) GetByteOrder() binary.ByteOrder {
	return binary.LittleEndian
}

type Unpack struct {
	ID   uint16
	Data []byte
}
