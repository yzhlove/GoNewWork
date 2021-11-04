package main

import (
	"fmt"
	"math/rand"
)

func main() {
	createAccount(20)
}

func createAccount(count int) {

	char := "abcdefghjkmnopqrstuvwxyz"
	size := len(char)

	reptUser := make(map[string]struct{})
	reptPasswd := make(map[string]struct{})

	realCount := 0
	for i := 0; i < count; i++ {
		idx := rand.Perm(size)

		var user, passwd string
		for j := 0; j < 4; j++ {
			user += fmt.Sprintf("%c", char[idx[j]])
			passwd += fmt.Sprintf("%c", char[idx[size-j-1]])
		}

		fmt.Println("username = ", user, " password = ", passwd)
		realCount++
		if _, ok := reptUser[user]; ok {
			i--
		} else {
			reptUser[user] = struct{}{}
		}

		if _, ok := reptPasswd[passwd]; ok {
			i--
		} else {
			reptPasswd[passwd] = struct{}{}
		}

	}

	fmt.Println("real count ", realCount)

}
