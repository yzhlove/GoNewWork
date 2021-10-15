package change

import "encoding/json"

type None struct {
	base
}

func (n *None) bind() {
	n.Typ = int(changedNone)
}

func (n None) Encode() ([]byte, error) {
	n.bind()
	return json.Marshal(n)
}

func (n None) decode(data []byte) (ChangedInterface, error) {
	var t = &None{}
	if err := json.Unmarshal(data, t); err != nil {
		return nil, err
	}
	return t, nil
}

func init() {
	register(int(changedNone), None{})
}
