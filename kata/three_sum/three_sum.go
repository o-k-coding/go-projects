package three_sum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)


	end := len(nums)

	for i := 0; i < end - 2; i++ {
		x := nums[i]
		// Skip if we have already used this value on the left to avoid left side duplicates
		if i > 0 && x == nums[i - 1] {
			continue
		}

		// Create two "Pointers" starting after the fixed left position and ending at the end of the slice
		j := i + 1
		k := end - 1

		// Loop until the pointers meet
		for j < k {
			y := nums[j]
			z := nums[k]

			// If the value on the left is the same as the previous value, skip it to avoid right side duplicates.
			// This could be more efficient to check only when decrementing... but I think this is a little more clear.
			// And avoids needing to set up another loop
			if k < end - 1 && z == nums[k + 1] {
				k --
				continue
			}

			sum := x + y + z

			if sum < 0 {
				// The left side is heavy, so increment the left pointer
				j ++
			} else if sum > 0 {
				// The right side is heavy so decrement the right pointer
				k --
			} else { // Sum is 0
				// Add the triplet
				result = append(result, []int{x,y,z})
				k --
			}
		}
	}

	return result
}
