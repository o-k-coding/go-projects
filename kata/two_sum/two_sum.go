package two_sum

func TwoSum(numbers []int, target int) [2]int {
	for i1 := 0; i1 < len(numbers); i1++ {
		for i2 := i1 + 1; i2 < len(numbers); i2++ {
			if (numbers[i1] + numbers[i2] == target) {
				return [2]int{i1, i2}
			}
		}
	}
  return [2]int{}
}
