package majority_element

// My original idea would be to use a hash map, but the problem asks to keep the space complexity O(1)
// so this works for that
func MajorityElement(nums []int) int {
	// Boyer-Moore Voting Algorithm
	// slect a candidate, then loop
	// If the number matches the candidate, add 1
	// otherwise subtract
	// if we hit 0, then select that number as the new candidate
	var candidate int
	count := 0

	for _, num := range nums{
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}


	return candidate
}
