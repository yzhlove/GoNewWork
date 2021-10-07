package main

import "testing"

func Test_Aes(t *testing.T) {

	ciphertext, nonce, err := Encoder([]byte("Hello World"), defineKey)
	if err != nil {
		t.Fatal(err)
	}

	text, err := Decoder(ciphertext, defineKey, nonce)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("text -> ", string(text))

}
