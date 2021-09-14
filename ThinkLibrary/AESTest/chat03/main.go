package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/aead/chacha20poly1305"
	"io"
	"os"
	"strings"
)

func main() {

	obtained := func(size int) []byte {
		data := make([]byte, size)
		io.ReadFull(rand.Reader, data)
		return data
	}

	nonce := obtained(12)
	key := obtained(32)

	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err :", err)
			return
		}

		str = strings.Replace(str, "\n", "", 1)
		fmt.Println("input string: ", str)

		if str == "stop" {
			return
		}

		encode := []byte(str)

		w, err := chacha20poly1305.EncryptWriter(bytes.NewBuffer(encode), key, nonce)
		if err != nil {
			fmt.Println("err :", err)
			return
		}
		w.Close()
		fmt.Println("encrpty string:", hex.EncodeToString(encode))

		decode := encode

		ww, err := chacha20poly1305.DecryptWriter(bytes.NewBuffer(decode), key, nonce)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		ww.Close()

		fmt.Println("decrpty string:", string(decode))

	}

}
