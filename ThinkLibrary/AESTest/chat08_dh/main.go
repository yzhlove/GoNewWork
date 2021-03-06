package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

/*
原理:
我们来看DH算法交换密钥的步骤。假设甲乙双方需要传递密钥，他们之间可以这么做：
甲首选选择一个素数p，例如509，底数g，任选，例如5，随机数a，例如123，然后计算A=g^a mod p，结果是215，然后，甲发送p＝509，g=5，A=215给乙；
乙方收到后，也选择一个随机数b，例如，456，然后计算B=g^b mod p，结果是181，乙再同时计算s=A^b mod p，结果是121；
乙把计算的B=181发给甲，甲计算s＝B^a mod p的余数，计算结果与乙算出的结果一样，都是121。
所以最终双方协商出的密钥s是121。注意到这个密钥s并没有在网络上传输。而通过网络传输的p，g，A和B是无法推算出s的，因为实际算法选择的素数是非常大的。
所以，更确切地说，DH算法是一个密钥协商算法，双方最终协商出一个共同的密钥，而这个密钥不会通过网络传输。
*/

// 添加密码本

var (
	// 2048-bit
	p2048 = new(big.Int).SetBytes(
		[]byte{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xC9, 0x0F, 0xDA, 0xA2, 0x21, 0x68, 0xC2, 0x34,
			0xC4, 0xC6, 0x62, 0x8B, 0x80, 0xDC, 0x1C, 0xD1,
			0x29, 0x02, 0x4E, 0x08, 0x8A, 0x67, 0xCC, 0x74,
			0x02, 0x0B, 0xBE, 0xA6, 0x3B, 0x13, 0x9B, 0x22,
			0x51, 0x4A, 0x08, 0x79, 0x8E, 0x34, 0x04, 0xDD,
			0xEF, 0x95, 0x19, 0xB3, 0xCD, 0x3A, 0x43, 0x1B,
			0x30, 0x2B, 0x0A, 0x6D, 0xF2, 0x5F, 0x14, 0x37,
			0x4F, 0xE1, 0x35, 0x6D, 0x6D, 0x51, 0xC2, 0x45,
			0xE4, 0x85, 0xB5, 0x76, 0x62, 0x5E, 0x7E, 0xC6,
			0xF4, 0x4C, 0x42, 0xE9, 0xA6, 0x37, 0xED, 0x6B,
			0x0B, 0xFF, 0x5C, 0xB6, 0xF4, 0x06, 0xB7, 0xED,
			0xEE, 0x38, 0x6B, 0xFB, 0x5A, 0x89, 0x9F, 0xA5,
			0xAE, 0x9F, 0x24, 0x11, 0x7C, 0x4B, 0x1F, 0xE6,
			0x49, 0x28, 0x66, 0x51, 0xEC, 0xE4, 0x5B, 0x3D,
			0xC2, 0x00, 0x7C, 0xB8, 0xA1, 0x63, 0xBF, 0x05,
			0x98, 0xDA, 0x48, 0x36, 0x1C, 0x55, 0xD3, 0x9A,
			0x69, 0x16, 0x3F, 0xA8, 0xFD, 0x24, 0xCF, 0x5F,
			0x83, 0x65, 0x5D, 0x23, 0xDC, 0xA3, 0xAD, 0x96,
			0x1C, 0x62, 0xF3, 0x56, 0x20, 0x85, 0x52, 0xBB,
			0x9E, 0xD5, 0x29, 0x07, 0x70, 0x96, 0x96, 0x6D,
			0x67, 0x0C, 0x35, 0x4E, 0x4A, 0xBC, 0x98, 0x04,
			0xF1, 0x74, 0x6C, 0x08, 0xCA, 0x18, 0x21, 0x7C,
			0x32, 0x90, 0x5E, 0x46, 0x2E, 0x36, 0xCE, 0x3B,
			0xE3, 0x9E, 0x77, 0x2C, 0x18, 0x0E, 0x86, 0x03,
			0x9B, 0x27, 0x83, 0xA2, 0xEC, 0x07, 0xA2, 0x8F,
			0xB5, 0xC5, 0x5D, 0xF0, 0x6F, 0x4C, 0x52, 0xC9,
			0xDE, 0x2B, 0xCB, 0xF6, 0x95, 0x58, 0x17, 0x18,
			0x39, 0x95, 0x49, 0x7C, 0xEA, 0x95, 0x6A, 0xE5,
			0x15, 0xD2, 0x26, 0x18, 0x98, 0xFA, 0x05, 0x10,
			0x15, 0x72, 0x8E, 0x5A, 0x8A, 0xAC, 0xAA, 0x68,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	// 768-bit
	p768 = new(big.Int).SetBytes(
		[]byte{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xC9, 0x0F, 0xDA, 0xA2, 0x21, 0x68, 0xC2, 0x34,
			0xC4, 0xC6, 0x62, 0x8B, 0x80, 0xDC, 0x1C, 0xD1,
			0x29, 0x02, 0x4E, 0x08, 0x8A, 0x67, 0xCC, 0x74,
			0x02, 0x0B, 0xBE, 0xA6, 0x3B, 0x13, 0x9B, 0x22,
			0x51, 0x4A, 0x08, 0x79, 0x8E, 0x34, 0x04, 0xDD,
			0xEF, 0x95, 0x19, 0xB3, 0xCD, 0x3A, 0x43, 0x1B,
			0x30, 0x2B, 0x0A, 0x6D, 0xF2, 0x5F, 0x14, 0x37,
			0x4F, 0xE1, 0x35, 0x6D, 0x6D, 0x51, 0xC2, 0x45,
			0xE4, 0x85, 0xB5, 0x76, 0x62, 0x5E, 0x7E, 0xC6,
			0xF4, 0x4C, 0x42, 0xE9, 0xA6, 0x3A, 0x36, 0x20,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		},
	)
	// p：素数 g：底数
	g    = big.NewInt(2)
	zero = big.NewInt(0)
)

// Sign 返回x的正负号。x<0时返回-1；x>0时返回+1；否则返回0。
// Cmp 比较x和y的大小。x<y时返回-1；x>y时返回+1；否则返回0。
// PubKey 范围 [1 , p-1]

type DhKey struct {
	PriKey *big.Int
	PubKey *big.Int
}

func ComputeKey(pubKey, priKey *big.Int) (*big.Int, error) {
	if pubKey == nil {
		return nil, errors.New("invalid pub key")
	}
	if pubKey.Sign() <= 0 && p768.Cmp(pubKey) <= 0 {
		return nil, errors.New("pub key out of bound")
	}
	if priKey == nil {
		return nil, errors.New("invalid pri key")
	}

	return big.NewInt(0).Exp(pubKey, priKey, p768), nil
}

func GenerateKey() (*DhKey, error) {
	var err error
	var x *big.Int

	for {
		if x, err = rand.Int(rand.Reader, p768); err != nil {
			return nil, err
		}

		if zero.Cmp(x) < 0 {
			break
		}
	}

	key := new(DhKey)
	key.PriKey = x
	key.PubKey = big.NewInt(0).Exp(g, x, p768)

	return key, nil
}

func main() {

	key1, err := GenerateKey()
	if err != nil {
		panic(err)
	}

	key2, err := GenerateKey()
	if err != nil {
		panic(err)
	}

	secret1, err := ComputeKey(key1.PubKey, key2.PriKey)
	if err != nil {
		panic(err)
	}

	secret2, err := ComputeKey(key2.PubKey, key1.PriKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("secret1:", secret1.String())
	fmt.Println("secret2:", secret2.String())

	if secret1.Cmp(secret2) == 0 {
		fmt.Println("succeed.")
	} else {
		fmt.Println("failed.")
	}

}
