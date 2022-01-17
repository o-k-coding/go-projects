package smallest_int_not_in_a

import "sort"

func SmallestIntNotInA(nums []int) int {
  // Find the smallest positive int that is NOT in nums
	smallestInt := 1
	if len(nums) == 0 {
		return smallestInt
	}
	sort.Ints(nums)

	for _, num := range nums {
		// Short circuit once the next number is greater than the current smallest
		if smallestInt < num {
			return smallestInt
		} else if smallestInt == num {
			smallestInt += 1
		}
	}
	return smallestInt
}
