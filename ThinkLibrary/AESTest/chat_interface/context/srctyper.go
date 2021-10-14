package context

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
)

type Field interface {
	Get() (key string, value interface{})
}

type String struct {
	key, value string
}

func WithString(key, value string) *String {
	return &String{key: key, value: value}
}

func (s *String) Get() (key string, value interface{}) {
	return s.key, s.value
}

type UInt struct {
	key   string
	value uint64
}

func WithUInt(key string, value uint64) *UInt {
	return &UInt{key: key, value: value}
}

func (u *UInt) Get() (key string, value interface{}) {
	return u.key, u.value
}

type Int struct {
	key   string
	value int64
}

func WithInt(key string, value int64) *Int {
	return &Int{key: key, value: value}
}

func (t *Int) Get() (key string, value interface{}) {
	return t.key, t.value
}

type Source struct {
	ItemSrcTyp int                      `json:"src_typ"` //道具来源
	Attach     map[string][]interface{} `json:"attach"`  //道具来源携带的附属信息
}

func (st *Source) Get() int {
	return st.ItemSrcTyp
}

func (st *Source) Encode() ([]byte, error) {
	return json.Marshal(st)
}

func (st *Source) Decode(data []byte) error {
	return json.Unmarshal(data, st)
}

func (st *Source) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("src_typ", st.ItemSrcTyp)
	return enc.AddReflected("attch", st.Attach)
}

func NewSource(typ int, fields ...Field) *Source {
	st := &Source{}
	st.ItemSrcTyp = typ
	st.Attach = make(map[string][]interface{}, len(fields))
	for _, t := range fields {
		k, v := t.Get()
		st.Attach[k] = append(st.Attach[k], v)
	}
	return st
}
