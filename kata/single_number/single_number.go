package single_number

func SingleNumber(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {

		if val, ok := counts[num]; ok {
			counts[num] = val +1
		} else {
			counts[num] = 1
		}
	}
	for k, v := range counts {
		if v == 1 {
			return k
		}
	}
	return 0
}


// func SingleNumber(nums []int) int {
// 	result := 0;
// 	for _, v := range nums {
// 		// Use xor bitwise operator, because repeatedly xoring the same value with itself will "flip" back to 0 or negated.
// 		// So whichever number isn't duplicated will not be negated
// 		result = result ^ v
// 		fmt.Println(v)
// 		fmt.Println(result)

// 	}
// 	return result;
// }
