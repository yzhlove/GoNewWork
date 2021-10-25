package pack

import (
	"bytes"
	"encoding/binary"
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
	x := make([]byte, packSize+msgSize+msg.Size())
	buf := bytes.NewBuffer(x)

	if err := binary.Write(buf, binary.LittleEndian, uint32(msgSize+msg.Size())); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, msgID); err != nil {
		return nil, err
	}

}
