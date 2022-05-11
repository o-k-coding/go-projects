package nth_tribonacci


func calculateTribonacci(n int, cache map[int]int) int {
	if cacheVal, ok := cache[n]; ok { return cacheVal }
	if n == 0 { return 0 }
	if n <= 2 { return 1 }
	cache[n] = calculateTribonacci(n - 1, cache) + calculateTribonacci(n - 2, cache) + calculateTribonacci(n - 3, cache)
	return cache[n]
}

func NthTribonacci(n int) int {
	return calculateTribonacci(n, make(map[int]int))
}
