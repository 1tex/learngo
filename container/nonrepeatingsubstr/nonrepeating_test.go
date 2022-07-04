package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是啊手动阀示范点", 10},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
