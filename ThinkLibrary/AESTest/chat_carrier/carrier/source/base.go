package source

import (
	"encoding/json"
)

type SrcInterface interface {
	bind() int
	Encode() ([]byte, error)
}

type decoder interface {
	decode(data []byte) (SrcInterface, error)
}

type base struct {
	Src int    `json:"src"`           //来源ID
	Ext string `json:"ext,omitempty"` //预留字段，额外需要记录的信息
}

func Decode(data string) (SrcInterface, error) {
	type a struct {
		Src int `json:"src"`
	}

	t := &a{}
	bytes := []byte(data)
	if err := json.Unmarshal(bytes, t); err != nil {
		return nil, err
	}

	if ret, ok := depot[t.Src]; ok {
		return ret.decode(bytes)
	}
	return nil, nil
}

var depot = make(map[int]decoder)

func register(src int, inter decoder) {
	depot[src] = inter
}
