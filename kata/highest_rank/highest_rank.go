package highest_rank

func HighestRank(nums []int) int {
	count := make(map[int]int)
	max := nums[0]

	for _, num := range nums {
		count[num] = count[num] + 1
		if max == num {
			continue
		} else if (count[max] < count[num] || (count[max] == count[num] && max < num)) {
			max = num
		}
	}
  return max
}
