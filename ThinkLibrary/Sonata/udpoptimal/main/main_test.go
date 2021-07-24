package main

import (
	"crypto/md5"
	"fmt"
	"net"
	"strings"
	"testing"
)

func Test_code(t *testing.T) {

	addr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 2379}

	encode := func(address string, index int) string {
		s := fmt.Sprintf("%s:%d", address, index)
		t.Log("sprintf:", s)
		ss := md5.Sum([]byte(s))
		t.Log("md5 code:", ss)
		ns := fmt.Sprintf("%s:%x", s, ss)
		t.Log("final string:", ns)
		return ns
	}

	endstr := encode(addr.String(), 1)

	t.Log(endstr)

	decode := func(str string, remote string, count int) {
		res := strings.Split(str, ":")
		s := fmt.Sprintf("%s:%d", remote, count)
		ss := md5.Sum([]byte(s))

		t.Logf("checkstr:%s gener:%s \n", res[len(res)-1], fmt.Sprintf("%x", ss))

		if res[len(res)-1] == fmt.Sprintf("%x", ss) {
			t.Log("success.")
		} else {
			t.Log("failed.")
		}
	}

	decode(endstr, addr.String(), 1)

}
