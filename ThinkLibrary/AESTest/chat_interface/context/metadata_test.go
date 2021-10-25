package context

import (
	"testing"
)

func Test_Metadata(t *testing.T) {

	metadata := Metadata{}

	src := NewSource(100, WithString("hello", "world"), WithInt("count", 100))

	items := &Items{}
	items.WithItem(Item{ID: 10001, Cur: 100, Change: 50})
	items.WithItem(Item{ID: 10002, Cur: 200, Change: -50})

	//mails := &Mails{}
	//metadata.With(src, mails)
	metadata.With(src, items)

}
