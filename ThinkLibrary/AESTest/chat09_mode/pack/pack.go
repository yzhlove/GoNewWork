package pack

import (
	"bytes"
	"encoding/binary"
)

const (
	PackSize = 4 // 包体长度
	MsgSize  = 2 // 消息头长度
)

type Msg interface {
	Size() int
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

//func Pack(msgID uint16, msg Msg) ([]byte, error) {
//	buf := bytes.NewBuffer([]byte{})
//
//	if err := binary.Write(buf, binary.LittleEndian, uint32(MsgSize+msg.Size())); err != nil {
//		return nil, err
//	}
//
//	if err := binary.Write(buf, binary.LittleEndian, msgID); err != nil {
//		return nil, err
//	}
//
//	data, err := msg.Marshal()
//	if err != nil {
//		return nil, err
//	}
//
//	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
//		return nil, err
//	}
//
//	return buf.Bytes(), nil
//}

func Pack(msgID uint16, data []byte) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, binary.LittleEndian, uint32(MsgSize+len(data))); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, msgID); err != nil {
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
	var data = make([]byte, len(p)-MsgSize)
	buf := bytes.NewReader(p[MsgSize:])

	if err := binary.Read(buf, binary.LittleEndian, &data); err != nil {
		return err
	}

	return msg.Unmarshal(data)
}
