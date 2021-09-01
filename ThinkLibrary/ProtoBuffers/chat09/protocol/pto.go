package protocol

//go:generate msgp -tests=false -io=false

type BInts []uint32

//msgp:ignore A
type A struct {
	B BInts
}



