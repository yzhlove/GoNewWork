package source

import "encoding/json"

type None struct {
	base
}

func (n *None) bind() {
	n.Src = int(srcNone)
}

func (n None) Encode() ([]byte, error) {
	n.bind()
	return json.Marshal(n)
}

func (n None) decode(data []byte) (SrcInterface, error) {
	var t = &None{}
	if err := json.Unmarshal(data, t); err != nil {
		return nil, err
	}
	return t, nil
}

func init() {
	register(int(srcNone), None{})
}
