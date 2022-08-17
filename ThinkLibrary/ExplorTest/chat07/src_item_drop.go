package main

type itemDrop struct {
	srcbase

	Act string `json:"act,omitempty"`
}

func (item itemDrop) Id() int {
	return int(srcItemDrop)
}

func (item itemDrop) Equal(src sourcer) bool {
	if ret, ok := src.(itemDrop); ok {
		if item.equal(ret.srcbase) {
			return item.Act == ret.Act
		}
	}
	return false
}
