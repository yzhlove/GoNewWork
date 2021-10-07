package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

var defineKey = []byte("*#06#-*#06#-*#06#-*#06#-*#06#-*#")

func Encoder(text, key []byte) (ciphertext, nonce []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce = getBytesWithSize(c.NonceSize())
	ciphertext = c.Seal(nonce, nonce, text, nil)

	return ciphertext, nonce, nil
}

func Decoder(ciphertext, key, nonce []byte) (original []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return c.Open(nil, ciphertext[:c.NonceSize()], ciphertext[c.NonceSize():], nil)
}

func getBytesWithSize(size int) []byte {
	ret := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, ret); err != nil {
		log.Fatal(err)
	}
	return ret
}
