package search_insert_pos

func SearchInsertPos(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i, num := range nums {
		if (num > target || num == target) {
			return i
		}
	}
	return len(nums)
}
