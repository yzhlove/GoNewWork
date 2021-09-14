package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//AES GCM加密解密
//secret 必须是16位或者32位 16:aes-128 32:aes-256

func main() {

	key := obtainedKey(16)

	data, nonce, err := encrypt([]byte("hello world"), key)
	if err != nil {
		panic(err)
	}

	ret, err := decrypt(data, key, nonce)
	if err != nil {
		panic(err)
	}

	fmt.Println("ret string : ", string(ret))
}

func encrypt(text, key []byte) ([]byte, []byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, c.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := c.Seal(nil, nonce, text, nil)
	fmt.Println("encrypt data:", hex.EncodeToString(ciphertext))

	return ciphertext, nonce, nil
}

func decrypt(data, key, nonce []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	source, err := c.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	return source, nil
}

func obtainedKey(size int) []byte {
	keys := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, keys); err != nil {
		panic("obtained key err:" + err.Error())
	}
	fmt.Println("obtained key string:", hex.EncodeToString(keys))
	return keys
}
