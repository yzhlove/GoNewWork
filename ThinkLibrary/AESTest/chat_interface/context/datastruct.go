package context

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
)

type Item struct {
	ID          uint32 //道具ID
	Cur, Change int32  //当前值，变化值
}

func (it Item) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint32("item_id", it.ID)
	enc.AddInt32("current", it.Cur)
	enc.AddInt32("change", it.Change)
	return nil
}

type itemArray []Item

func (ss itemArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range ss {
		if err := arr.AppendObject(ss[i]); err != nil {
			return err
		}
	}
	return nil
}

type Items struct {
	Data itemArray
}

func (it *Items) WithItem(t Item) {
	it.Data = append(it.Data, t)
}

func (it Items) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return enc.AddArray("items_change", it.Data)
}

func (it Items) Encode() ([]byte, error) {
	return json.Marshal(it)
}

func (it *Items) Decode(data []byte) error {
	return json.Unmarshal(data, it)
}
