package b

type Inter interface {
	Init(bb *B)
	Run() string
}

type B struct {
	inter Inter
}

func (b *B) Obt() Inter {
	return b.inter
}

func (b *B) Init(inter Inter) {
	b.inter = inter
}
