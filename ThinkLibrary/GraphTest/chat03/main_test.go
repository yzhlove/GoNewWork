package main

import "testing"

func Test_Update(t *testing.T) {

	cases := []struct {
		name      string
		cur_star  int
		materials []material
		upgrade   int
		refund    int
	}{
		{"A", 0, []material{{1, 2}}, 2, 0},
		{"B", 0, []material{{1, 3}}, 2, 1},
		{"C", 0, []material{{1, 4}}, 3, 1},
		{"D", 0, []material{{1, 10}}, 5, 4},
	}

	p := presents{}
	p.load()
	p.show()

	for _, c := range cases {
		a, b, err := p.Update(c.cur_star, c.materials)
		if err != nil {
			t.Logf("[%s] case error, [%v]", c.name, err)
			return
		}

		if c.upgrade != a && c.refund != b.number {
			t.Logf("[%s] case failed, [%d] [%d] expected-> [%d] [%d]", c.name, a, b.number, c.upgrade, c.refund)
		}
	}

	t.Log("Ok.")
}
