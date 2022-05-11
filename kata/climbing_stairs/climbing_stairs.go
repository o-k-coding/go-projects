package climbing_stairs

// n steps to the top of the stairs
// You can climb 1 or 2 steps at a time
// How many distinct ways can you climb to the top?
// Think of the problem in a tree structure where from each step you have two children, 1 where you take 1 step and 1 where you take 2
// If you can no longer take 1 or 2 steps you are done. If you made it to the top, keep that number!

func calculatePaths(currentStair int, top int, cache map[int]int) int {
	if pathCount, ok := cache[currentStair]; ok {
		return pathCount
	}
	if (currentStair > top) {
		return 0
	}
	if (currentStair == top) {
		return 1
	}
	cache[currentStair] = calculatePaths(currentStair + 1, top, cache) + calculatePaths(currentStair + 2, top, cache)
	return cache[currentStair]
}
func ClimbingStairs(n int) int {
	return calculatePaths(0, n, make(map[int]int))
}
