package secret

import (
	"testing"
	"think-library/AESTest/chat09_mode/conf"
)

func Test_Aes(t *testing.T) {

	a, err := NewAes(conf.DefineKey)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := a.Encode([]byte("hello world"))
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("encoder = %x", data)

	text, err := a.Decode(data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("decoder = %s", text)

}
