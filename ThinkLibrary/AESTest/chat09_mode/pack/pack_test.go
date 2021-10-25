package pack

import (
	"testing"
	"think-library/AESTest/chat09_mode/pb"
)

func Test_Pack(t *testing.T) {

	var msgId uint16 = 1000
	auth := &pb.Auth{Msg: "dh algorithm", PubKey: []byte("Public Key")}

	data, err := Pack(msgId, auth)
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
