package context

import "testing"

func Test_SrcType(t *testing.T) {

	st := NewSource(100,
		WithString("a", "aa"),
		WithString("a", "ab"),
		WithString("b", "bbb"),
		WithInt("aInt", 123),
		WithUInt("aInt", 456),
		WithInt("bInt", 112233))

	data, err := st.Encode()
	if err != nil {
		t.Error("encoder -> ", err)
		return
	}

	t.Log("data string =>", string(data))

}
