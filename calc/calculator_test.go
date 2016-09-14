package calc

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		ain, bin, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{100, -110, -10},
	}

	for _, c := range cases {
		got := Add(c.ain, c.bin)
		if got != c.want {
			t.Errorf("Add(%d, %d) == %d, want %d", c.ain, c.bin, got, c.want)
		}
	}
}
