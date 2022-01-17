package race_results

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func ParseTimeIntoSeconds(time string) int {
	parsedTime := strings.Split(time, "|")
	if len(parsedTime) != 3 {
		return 0
	}
	hours, err := strconv.Atoi(parsedTime[0])
	if err != nil {
		hours = 0
	}
	minutes, err := strconv.Atoi(parsedTime[1])
	if err != nil {
		minutes = 0
	}
	secondsStr := parsedTime[2]
	if len(secondsStr) > 2 {
		secondsStr = secondsStr[0:2]
	}
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		seconds = 0
	}
	return seconds + (minutes * 60) + (hours * 60 * 60)
}

func ParseSecondsIntoTime(totalSeconds int) string {
	secondsF := float64(totalSeconds)
	hours := math.Floor(secondsF / (60 * 60))
	// Remainder of hours calculation / 60
	hoursRemainder := totalSeconds % (60 * 60)
	minutes := math.Floor(float64(hoursRemainder) / 60)
	// Remainder of minutes calculation / 60
	seconds := hoursRemainder % 60
	// %02d left pads the number value with 0s to 2 digits
	return fmt.Sprintf("%02d|%02d|%02d", int(hours), int(minutes), seconds)
}

func CalculateRange(min int, max int) string {
	return ParseSecondsIntoTime(max - min)
}

func CalculateAverage(sum int, total int) string {
	return ParseSecondsIntoTime(sum / total)
}

func CalculateMedian(times []int) string {
	length := len(times)

	if length%2 == 0 { // Even number of times
		half := length / 2
		value1 := times[half-1]
		value2 := times[half]

		median := (value1 + value2) / 2
		return ParseSecondsIntoTime(median)
	} else { // Odd number of times
		middle := int(math.Ceil(float64(length)/2)) - 1
		return ParseSecondsIntoTime(times[middle])
	}
}

func Stati(results string) string {
	if len(results) == 0 {
		return ""
	}
	teamResults := strings.Split(results, ",")
	resultTimes := []int{}
	sum := 0

	for _, result := range teamResults {
		resultCleaned := strings.TrimSpace(result)
		resultTime := ParseTimeIntoSeconds(resultCleaned)
		resultTimes = append(resultTimes, resultTime)
		sum += resultTime
	}

	sort.Ints(resultTimes)
	rangeResult := CalculateRange(resultTimes[0], resultTimes[len(resultTimes) - 1])
	averageResult := CalculateAverage(sum, len(resultTimes))
	medianResult := CalculateMedian(resultTimes)

	return fmt.Sprintf("Range: %s Average: %s Median: %s",
		rangeResult,
		averageResult,
		medianResult,
	)
}
