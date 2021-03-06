package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/aead/chacha20poly1305"
	"io"
	"log"
)

func main() {

	key := []byte("*#06#-*#06#-*#06#-*#06#-*#06#-*#")

	//a, err := NewAes(key)
	a, err := NewChacha20(key)
	if err != nil {
		panic(err)
	}

	data, err := a.Encode([]byte("hello world"))
	if err != nil {
		panic(err)
	}

	src, err := a.Decode(data)
	if err != nil {
		panic(err)
	}

	log.Print("src -> ", string(src))

}

type Aes struct {
	key  []byte
	aead cipher.AEAD
}

func NewAes(key []byte) (*Aes, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	c, err := cipher.NewGCM(b)
	if err != nil {
		return nil, err
	}
	return &Aes{key: key, aead: c}, nil
}

func NewChacha20(key []byte) (*Aes, error) {
	aead, err := chacha20poly1305.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &Aes{key: key, aead: aead}, nil
}

func (a *Aes) Encode(text []byte) ([]byte, error) {

	nonce, err := a.getNonce()
	if err != nil {
		return nil, err
	}

	ciphertext := a.aead.Seal(nonce, nonce, text, nil)
	return ciphertext, nil
}

func (a *Aes) Decode(ciphertext []byte) ([]byte, error) {
	size := a.aead.NonceSize()
	return a.aead.Open(nil, ciphertext[:size], ciphertext[size:], nil)
}

func (a *Aes) getNonce() ([]byte, error) {
	nonce := make([]byte, a.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}
