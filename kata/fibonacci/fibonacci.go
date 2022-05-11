package fibonacci

func calculateFibonacci(currentN int, cache map[int]int) int {
		if cacheValue, ok := cache[currentN]; ok { return cacheValue }
		if currentN == 0 { return 0 }
		if currentN == 1 { return 1 }
		cache[currentN] = calculateFibonacci(currentN - 1, cache) + calculateFibonacci(currentN - 2, cache)
		return cache[currentN]
	}

func Fibonacci(n int) int {
	cache := make(map[int]int)
	return calculateFibonacci(n, cache)
}
