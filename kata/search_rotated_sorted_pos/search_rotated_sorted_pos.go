package search_rotated_sorted_pos

func SearchRotatedSortedPos(nums []int, target int) int {
	if (target == nums[0]) {
		return 0
	}
	modifier := 1
	i := 0
	end := len(nums)
	if (target < nums[0]) {
		modifier = -1
		i = len(nums) - 1
		end = 0
	}


	for i != end {
		if (nums[i] == target) {
			return i
		} else if (modifier == -1 && nums[i] < target) {
			break
		} else if (modifier == 1 && nums[i] > target) {
			break
		}
		i += 1 * modifier
	}

	return -1
}
