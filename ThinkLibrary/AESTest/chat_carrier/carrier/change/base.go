package change

import "encoding/json"

type ChangedInterface interface {
	bind()
	Encode() ([]byte, error)
}

type decoder interface {
	decode(data []byte) (ChangedInterface, error)
}

type base struct {
	Typ int `json:"typ"`
}

func Decode(data string) (ChangedInterface, error) {
	type a struct {
		Typ int `json:"typ"`
	}

	t := &a{}
	bytes := []byte(data)
	if err := json.Unmarshal(bytes, t); err != nil {
		return nil, err
	}

	if ret, ok := depot[t.Typ]; ok {
		return ret.decode(bytes)
	}
	return nil, nil
}

var depot = make(map[int]decoder)

func register(typ int, inter decoder) {
	depot[typ] = inter
}
