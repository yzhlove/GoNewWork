package pack

import (
	"bytes"
	"encoding/binary"
	"io"
	"unsafe"
)

type Message struct {
	conn io.ReadWriter
	pack MessageInterface
}

func (m *Message) Encode() error {

	if m.conn == nil {
		return nil
	}

	dataLen := int(unsafe.Sizeof(m.pack.ID())) + len(m.pack.Data())
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, m.pack.ByteOrder(), uint32(dataLen)); err != nil {
		return err
	}

	if err := binary.Write(buf, m.pack.ByteOrder(), m.pack.ID()); err != nil {
		return err
	}

	if err := binary.Write(buf, m.pack.ByteOrder(), m.pack.Data()); err != nil {
		return err
	}

	if _, err := m.conn.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
