package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(uint64(i))
	}
}

func TestPopCountLoop(t *testing.T) {
	for i := 0; i < 100; i++ {
		if e, r := PopCount(uint64(i)), PopCountLoop(uint64(i)); e != r {
			t.Errorf("PopCountLoop(%v) = %v, wanted %v", i, r, e)
		}
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(uint64(i))
	}
}

func TestPopCountShift(t *testing.T) {
	for i := 0; i < 100; i++ {
		if e, r := PopCount(uint64(i)), PopCountShift(uint64(i)); e != r {
			t.Errorf("PopCountShift(%v) = %v, wanted %v", i, r, e)
		}
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(uint64(i))
	}
}

func TestPopCountClear(t *testing.T) {
	for i := 0; i < 100; i++ {
		if e, r := PopCount(uint64(i)), PopCountClear(uint64(i)); e != r {
			t.Errorf("PopCountClear(%v) = %v, wanted %v", i, r, e)
		}
	}
}
