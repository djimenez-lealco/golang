package main

import "testing"

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func BenchmarkAbs(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Abs(i)
	}
}

func FuzzAbs(f *testing.F) {
	f.Add(-1)
	f.Add(-2)
	f.Fuzz(func(t *testing.T, in int) {
		out := Abs(in)
		if out != in*-1 {
			t.Errorf("Abs(%d) != %d", in, out)
		}
	})
}
