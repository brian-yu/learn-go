package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var c byte
	var i uint
	for i = 0; i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}

func PopCountShift(x uint64) int {
	c := 0
	for i := 0; i < 64; i++ {
		c += int(x & 1)
		x = x >> 1
	}
	return c
}

func PopCountClear(x uint64) (c int) {
	for x != 0 {
		c++
		x = x & (x - 1)
	}
	return
}
