package source

import "encoding/json"

type Items struct {
	base
}

func (it *Items) bind() int {
	return int(srcItems)
}

func (it Items) Equal(src Items) bool {
	if it.
}

func (it *Items) Encode() ([]byte, error) {
	it.Src = it.bind()
	return json.Marshal(it)
}

func (it Items) decode(data []byte) (SrcInterface, error) {
	var t = &Items{}
	if err := json.Unmarshal(data, t); err != nil {
		return nil, err
	}
	return t, nil
}

func init() {
	register(int(srcItems), Items{})
}
