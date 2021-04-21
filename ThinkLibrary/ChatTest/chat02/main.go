package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

//ras test

func main() {
	test2()
}

func test1() {
	pubKey, privKey := GenRsaKey()
	fmt.Println(string(pubKey))
	fmt.Println(string(privKey))

	var data = []byte("abcdefghijklmn")
	fmt.Println("对消息进行签名操作...")
	signData := RsaSingWithSha256(data, privKey)
	fmt.Println("消息的签名信息:->", hex.EncodeToString(signData))
	fmt.Println()
	fmt.Println("对签名的消息进行验证...")
	if ResVerifySignWithSha256(data, signData, pubKey) {
		fmt.Println("签名信息验证成功，确定是正确的签名!!!")
	}
	fmt.Println(strings.Repeat("-", 50))
	ciphertext := RsaEncrypt(data, pubKey)
	fmt.Println("公钥加密后的数据:", hex.EncodeToString(ciphertext))
	source := RsaDecrypt(ciphertext, privKey)
	fmt.Println("私钥解密后的数据:", string(source))

}

//公钥私钥不能颠倒,mmp🙃️
func test2() {
	pubKey, privKey := GenRsaKey()

	var data = []byte("老子天下第一!!!")
	fmt.Println("对消息进行签名...")
	sign := RsaSingWithSha256(data, privKey)
	fmt.Println("->", hex.EncodeToString(sign))
	fmt.Println()
	fmt.Println("对签名的消息进行验证...")
	if ResVerifySignWithSha256(data, sign, pubKey) {
		fmt.Println("签名的消息验证成功!")
	}

	fmt.Println(strings.Repeat("=", 50))
	cipher := RsaEncrypt(data, pubKey)
	fmt.Println("使用私钥加密后的数据->", hex.EncodeToString(cipher))
	src := RsaDecrypt(cipher, privKey)
	fmt.Println("使用公钥解密后的数据->", string(src))
}

func GenRsaKey() (pubkey, prvkey []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes}
	prvkey = pem.EncodeToMemory(block)

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	block = &pem.Block{Type: "PUBLIC KEY", Bytes: derPkix}
	pubkey = pem.EncodeToMemory(block)
	return
}

func RsaSingWithSha256(data, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}

	if privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		panic(fmt.Errorf("ParsePKCS1PrivateKey error:%s", err))
	} else {
		if signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed); err != nil {
			panic(fmt.Errorf("error form signing:%s", err))
		} else {
			return signature
		}
	}
	return nil

}

func ResVerifySignWithSha256(data, signDate, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}

	if publicKey, err := x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		panic(err)
	} else {
		hashed := sha256.Sum256(data)
		if err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signDate); err != nil {
			panic(err)
		}
		return true
	}

	return false
}

func RsaEncrypt(data, keyBytes []byte) []byte {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	if pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		panic(err)
	} else {
		pub := pubInterface.(*rsa.PublicKey)
		if ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data); err != nil {
			panic(err)
		} else {
			return ciphertext
		}
	}
	return nil
}

func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	if priv, err := x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		panic(err)
	} else {
		if data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext); err != nil {
			panic(err)
		} else {
			return data
		}
	}
	return nil
}
