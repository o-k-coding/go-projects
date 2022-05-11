package can_place_flowers

func CanPlaceFlowers(flowerbed []int, n int) bool {
	// 0 means no flower
	// 1 means flower
	// Flowers cannot be placed in adjacent pots
	// Can you place n number of new flowers?

	// First thought, calculate the number of 0s that are not adjacent to a 1 in the flowerbed
	validSpaces := 0
	for i, flower := range flowerbed {
		if flower == 0 {
			leftFlowerOpen := i == 0 || flowerbed[i - 1] == 0
			if (!leftFlowerOpen) {continue}
			rightFlowerOpen := (i == len(flowerbed) - 1) || flowerbed[i + 1] == 0
			if (!rightFlowerOpen) {continue}
			if leftFlowerOpen && rightFlowerOpen {
				validSpaces ++
				flowerbed[i] = 1
			}
		}
	}

	return n <= validSpaces
}
