package min_cost_climbing_stairs

// n steps to the top of the stairs
// You can climb 1 or 2 steps at a time
// How many distinct ways can you climb to the top?
// Think of the problem in a tree structure where from each step you have two children, 1 where you take 1 step and 1 where you take 2
// If you can no longer take 1 or 2 steps you are done. If you made it to the top, keep that number!

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// cache will be a
func calculatePathCost(currentStair int, top int, cost []int, costCache map[int]int) int {
	if pathCost, ok := costCache[currentStair]; ok {
		return pathCost
	}
	if (currentStair >= top) {
		return 0
	}
	oneStepCost := calculatePathCost(currentStair + 1, top, cost, costCache)
	twoStepCost := calculatePathCost(currentStair + 2, top, cost, costCache)
	costCache[currentStair] = cost[currentStair] + min(oneStepCost, twoStepCost)
	return costCache[currentStair]
}
func MinCostClimbingStairs(cost []int) int {
	// Problems gives the assumption that there are at least 2 stairs
	top := len(cost)
	zeroStart := calculatePathCost(0, top, cost, make(map[int]int))
	oneStart := calculatePathCost(1, top, cost, make(map[int]int))
	return min(zeroStart, oneStart)
}
