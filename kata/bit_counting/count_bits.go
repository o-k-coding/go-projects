package bit_counting

import (
	"strconv"
	"strings"
)

func CountBits(num uint) int {
	binaryNum := strconv.FormatInt(int64(num), 2)
	onesCount := strings.Count(binaryNum, "1")
	return onesCount
}
