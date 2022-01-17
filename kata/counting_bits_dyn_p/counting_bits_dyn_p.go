package counting_bits_dyn_p

import (
	"math"
)

func OneCountBinary(num int, cache map[int] int) int {
	if num == 0 {
		return 0
	}
	if val, ok := cache[num]; ok {
		return val
	}
	bin := num % 2
	quotient := int(math.Floor(float64(num) / 2))
	cache[num] = bin + OneCountBinary(quotient, cache)
	return cache[num]
}

func CountingBitsDynP(num int) []int {
	result := []int{0}
	cache := make(map[int]int)
	for i := 1; i <= num; i++ {
		result = append(result, OneCountBinary(i, cache))
	}
	return result
}
