package main

import "fmt"

func main() {

	var s SIGN
	s.Update(1000)
	fmt.Println(s.Get())

}

type SIGN uint32

func (s *SIGN) Update(n uint32) {
	*s = SIGN(n)
}

func (s SIGN) Get() uint32 {
	return uint32(s)
}
