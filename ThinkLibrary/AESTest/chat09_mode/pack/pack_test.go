package pack

import (
	"bytes"
	"encoding/binary"
	"testing"
	"think-library/AESTest/chat09_mode/pb"
)

func Test_Pack(t *testing.T) {

	var msgId uint16 = 1000
	auth := &pb.Auth{Msg: "dh algorithm", PubKey: []byte("Public Key")}

	data, err := auth.Marshal()
	if err != nil {
		t.Error(err)
		return
	}

	data, err = Pack(msgId, data)
	if err != nil {
		t.Error(err)
		return
	}

	// 前四个字节是为了防止Tcp粘包而填入的消息头大小
	tbl := Packet(data[4:])
	t.Log("msgId:", tbl.Id())

	auth2 := &pb.Auth{}
	if err := tbl.Unpack(auth2); err != nil {
		t.Error(err)
		return
	}

	t.Log("msg:", auth2)

}

func Test_Parse(t *testing.T) {

	data := []byte("hello world")

	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		t.Error(err)
		return
	}

	t.Log("buf => ", string(buf.Bytes()))

	decode := make([]byte, buf.Len())
	if err := binary.Read(buf, binary.LittleEndian, &decode); err != nil {
		t.Error(err)
		return
	}

	t.Log("decode string => ", string(decode))

}
