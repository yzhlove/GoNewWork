package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/aead/chacha20poly1305"
	"io"
)

// key size must 32 bits

func main() {

	key := obtainedBytes(32)
	nonce := obtainedBytes(8)

	data, err := cipherEncrpty([]byte("hello "), nonce, key)
	if err != nil {
		panic(err)
	}

	if _, err := cipherDecrpty(data, nonce, key); err != nil {
		panic(err)
	}

	nonce2 := obtainedBytes(12)

	data, err = cipherIETFEncrpty([]byte("world "), nonce2, key)
	if err != nil {
		panic(err)
	}

	if _, err := cipherIETFDecrpty(data, nonce2, key); err != nil {
		panic(err)
	}

}

func cipherEncrpty(source, nonce, key []byte) ([]byte, error) {

	aead, err := chacha20poly1305.NewCipher(key)
	if err != nil {
		return nil, err
	}

	fmt.Println("nonce size -> ", aead.NonceSize())

	data := aead.Seal(nil, nonce, source, nil)
	fmt.Println("encrpty data string: ", hex.EncodeToString(data))
	return data, nil
}

func cipherIETFEncrpty(text, nonce, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewIETFCipher(key)
	if err != nil {
		return nil, err
	}

	fmt.Println("nonce size -> ", aead.NonceSize())

	ret := aead.Seal(nil, nonce, text, nil)
	fmt.Println("ietf encrpty data string: ", hex.EncodeToString(ret))
	return ret, nil
}

func cipherDecrpty(data, nonce, key []byte) ([]byte, error) {

	aead, err := chacha20poly1305.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ret, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("decrpty data string: ", string(ret))

	return ret, nil
}

func cipherIETFDecrpty(data, nonce, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewIETFCipher(key)
	if err != nil {
		return nil, err
	}

	ret, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("ietf decrpty data string: ", string(ret))

	return ret, nil
}

func obtainedBytes(size int) []byte {
	data := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		panic(err)
	}
	return data
}
