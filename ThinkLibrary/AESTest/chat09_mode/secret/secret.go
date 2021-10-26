package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type EncoderDecoder interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

type AesSecret struct {
	aead cipher.AEAD
}

func (a *AesSecret) Encode(data []byte) ([]byte, error) {

	nonce := make([]byte, a.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return a.aead.Seal(nonce, nonce, data, nil), nil
}

func (a *AesSecret) Decode(data []byte) ([]byte, error) {
	s := a.aead.NonceSize()
	return a.aead.Open(nil, data[:s], data[s:], nil)
}

func NewAes(key []byte) (*AesSecret, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &AesSecret{aead: aead}, nil
}
