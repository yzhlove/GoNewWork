package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"sync"
	"time"
)

const HOST = "127.0.0.1:50051"

func main() {

	go func() {
		if err := server(); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Millisecond)
	if err := client(); err != nil {
		log.Fatal(err)
	}
}

func appendKey(data []byte) []byte {
	ret := make([]byte, 32)
	copy(ret, data)
	return ret
}

func client() error {
	conn, err := net.Dial("tcp", HOST)
	if err != nil {
		return err
	}

	msg := Msg{conn: conn}

	//1.客户端拿到自己的private,public两个big.Int
	C2, E2 := Exp()

	//2.使用AES加密public
	ciphertext, nonce, err := Encoder(E2.Bytes(), defineKey)
	if err != nil {
		return err
	}
	//3.向服务器发送 Aes(public) + nonce
	if err := msg.Write(SystemConnect, ciphertext); err != nil {
		return err
	}

	//3.向服务器发送nonce
	if err := msg.Write(SystemNonce, nonce); err != nil {
		return err
	}

	//4.等待服务器发送自己的public
	id, data, err := msg.Read()
	if err != nil {
		return err
	}

	//检查ID是否正确
	var sevPubData []byte
	if id == SystemConnect {
		sevPubData = make([]byte, len(data))
		copy(sevPubData, data)
	} else {
		return errors.New("replay message id error")
	}

	id, data, err = msg.Read()
	if err != nil {
		return err
	}

	var aesKey []byte
	if id == SystemNonce {
		E1, err := Decoder(sevPubData, defineKey, data)
		if err != nil {
			return err
		}

		keyInt := Key(C2, big.NewInt(0).SetBytes(E1))
		aesKey = appendKey(keyInt.Bytes())
	}

	fmt.Println("aesKey ==> ", hex.EncodeToString(aesKey), " aesKey length => ", len(aesKey))

	go func() {
		var userData []byte
		for {
			tag, ret, erro := msg.Read()
			if erro != nil {
				log.Fatal("client read message error:", err)
			}

			switch tag {
			case UsersConnect:
				userData = make([]byte, len(ret))
				copy(userData, ret)
			case UsersNonce:
				text, erro := Decoder(userData, aesKey, ret)
				if erro != nil {
					log.Fatal("Client Decoder error:", erro)
				}
				log.Println("client:", string(text))
			}
		}
	}()

	var count = 1
	for {

		str := []byte(fmt.Sprintf("client send server message:%d", count))
		ciphertext, nonce, err := Encoder(str, aesKey)
		if err != nil {
			return err
		}

		if err := msg.Write(UsersConnect, ciphertext); err != nil {
			return err
		}

		if err := msg.Write(UsersNonce, nonce); err != nil {
			return err
		}
		count += 2

		time.Sleep(time.Second)
	}

	//return nil
}

func server() error {
	l, err := net.Listen("tcp", HOST)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print("accept error:", err)
			break
		}
		wg.Add(1)
		go serverHandle(&wg, conn)
	}
	wg.Wait()
	return nil
}

func serverHandle(wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()

	msg := Msg{conn: conn}

	S1, E1 := Exp()

	var cliPubData []byte
	var aesKey []byte

	var userData []byte
	var count = 0
	for {
		id, data, err := msg.Read()
		if err != nil {
			log.Fatal(err)
		}

		switch id {
		case SystemConnect:
			cliPubData = make([]byte, len(data))
			copy(cliPubData, data)
		case SystemNonce:
			E2, err := Decoder(cliPubData, defineKey, data)
			if err != nil {
				log.Fatal("Decoder -> ", err)
			}
			//新的AES KEY
			keyInt := Key(S1, big.NewInt(0).SetBytes(E2))

			//服务端加密自己的公钥
			ciphertext, nonce, err := Encoder(E1.Bytes(), defineKey)
			if err != nil {
				log.Fatal("Encoder -> ", err)
			}
			//向客户端发送自己加密后的公钥
			if err := msg.Write(SystemConnect, ciphertext); err != nil {
				log.Fatal("msg.Write SystemConnect-> ", err)
			}
			//向客户端发送自己的nonce
			if err := msg.Write(SystemNonce, nonce); err != nil {
				log.Fatal("msg.Write System Nonce-> ", err)
			}

			//设置新的AES key
			aesKey = appendKey(keyInt.Bytes())
			fmt.Println("aesKey ==> ", hex.EncodeToString(aesKey), " aesKey length => ", len(aesKey))

		case UsersConnect:
			userData = make([]byte, len(data))
			copy(userData, data)
		case UsersNonce:
			count += 2
			text, err := Decoder(userData, aesKey, data)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("server:", string(text))

			str := []byte(fmt.Sprintf("server send client message:%d", count))
			ciphertext, nonce, err := Encoder(str, aesKey)
			if err != nil {
				log.Fatal("server Encoder ", err)
			}

			if err := msg.Write(UsersConnect, ciphertext); err != nil {
				log.Fatal("server Write UserConnect ", err)
			}

			if err := msg.Write(UsersNonce, nonce); err != nil {
				log.Fatal("server Write UserNonce ", err)
			}
		}

	}

}
