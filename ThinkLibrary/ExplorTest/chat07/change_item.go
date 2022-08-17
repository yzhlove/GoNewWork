package main

type item struct {
	changebase

	UID    uint64 `json:"uid,omitempty"`
	Tid    uint32 `json:"tid"`
	Cur    int32  `json:"cur,omitempty"`
	Change int32  `json:"change,omitempty"`
}

func (it item) Id() int {
	return int(changeItem)
}
