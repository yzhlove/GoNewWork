package main

import "testing"

func Test_DH(t *testing.T) {

	X1, E1 := Exp()
	X2, E2 := Exp()

	keyInt1 := Key(X1, E2)
	keyInt2 := Key(X2, E1)

	if keyInt1.Cmp(keyInt2) == 0 {
		t.Log("success")
	} else {
		t.Log("failed")
	}

}
