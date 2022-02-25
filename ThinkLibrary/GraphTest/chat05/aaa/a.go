package aaa

//go:generate /Users/yostar/Develop/Go/GoPath/bin/msgp -tests=false -io=false

// msgp 不支持小写的结构体

type Base struct {
	at string
	bt string
	ct int
}

//msgp:ignore stc
type stc struct {
	Base

	str string
}

func (s *stc) marsh() ([]byte, error) {

	return nil, nil
}

func (s *stc) unmarsh(data []byte) error {

	return nil
}
