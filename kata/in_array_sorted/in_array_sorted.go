package in_array_sorted

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	results := []string{}
	resultsSet := make(map[string]bool)

	if len(array1) == 0 || len(array2) == 0 {
		return results
	}

	// First loop over array 1
	// Second loop over array 2
	// Compare array 1 item to each item in array 2 to determine if array 1 item is a substring of item in array 2
	// If a substring is found, add the array 1 item to the results arrray and break the loop

	for _, item1 := range array1 {
		// If we already have found a match, or there are duplicates for this item, move on
		if resultsSet[item1] {
			continue
		}
		for _, item2 := range array2 {
			if strings.Contains(item2, item1) {
				results = append(results, item1)
				resultsSet[item1] = true
				break
			}
		}
	}

	sort.Strings(results)
	return results // sorted array of strings in array1 which as substrings of an item in array2
}
