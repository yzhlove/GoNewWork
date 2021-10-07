package main

import (
	"bytes"
	"testing"
)

func Test_Pack(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	msg := Msg{conn: buf}
	if err := msg.Write(SystemConnect, []byte("Hello World")); err != nil {
		t.Fatal(err)
	}

	id, data, err := msg.Read()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("id:%d data:%v", id, string(data))
}
