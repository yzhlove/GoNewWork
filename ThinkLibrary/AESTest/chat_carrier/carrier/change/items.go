package change

import "encoding/json"

type Item struct {
	ID     uint32 `json:"id,omitempty"`     //道具ID
	Cur    int32  `json:"cur,omitempty"`    //当前值
	Change int32  `json:"change,omitempty"` //变化值
}

type Items struct {
	base
	Data []Item `json:"data"`
}

func (it *Items) Append(t Item) {
	it.Data = append(it.Data, t)
}

func (it *Items) bind() {
	it.Typ = int(changeItems)
}

func (it *Items) Encode() ([]byte, error) {
	it.bind()
	return json.Marshal(it)
}

func (it Items) decode(data []byte) (ChangedInterface, error) {
	var t = &Items{}
	if err := json.Unmarshal(data, t); err != nil {
		return nil, err
	}
	return t, nil
}

func init() {
	register(int(changeItems), Items{})
}
