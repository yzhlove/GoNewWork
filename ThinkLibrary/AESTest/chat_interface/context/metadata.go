package context

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
)

type logDataInterfaceArray []LogDataInterface

func (ss logDataInterfaceArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range ss {
		if err := arr.AppendObject(ss[i]); err != nil {
			return err
		}
	}
	return nil
}

type Adapter struct {
	Source LogSourceInterface    `json:"source"`
	List   logDataInterfaceArray `json:"list"`
}

func (c Adapter) Encode() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Adapter) Decode(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c Adapter) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if err := enc.AddObject("source", c.Source); err != nil {
		return err
	}
	return enc.AddArray("list", c.List)
}

type Metadata map[int]*Adapter

func (m Metadata) With(src LogSourceInterface, change LogDataInterface) {
	if _, ok := m[src.Get()]; !ok {
		m[src.Get()] = &Adapter{Source: src}
	}
	m[src.Get()].List = append(m[src.Get()].List, change)
}

func (m *Metadata) Clear() {
	*m = make(map[int]*Adapter)
}
