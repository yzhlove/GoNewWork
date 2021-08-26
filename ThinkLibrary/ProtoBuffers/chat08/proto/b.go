package proto

//go:generate msgp -tests=false -io=false

type AItem struct {
	Id  uint32
	Qty uint32
}

type AItems struct {
	List []AItem
}

type BItem struct {
	Id  uint32 `msg:"a,omitempty"`
	Qty uint32 `msg:"b,omitempty"`
}

type BItems struct {
	List []BItem `msg:"c,omitempty"`
}

//msgp:tuple CItem
type CItem struct {
	Id  uint32
	Qty uint32
}

type CItems struct {
	List []CItem `msg:"c,omitempty"`
}
