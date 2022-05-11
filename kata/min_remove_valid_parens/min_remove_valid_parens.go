package min_remove_valid_parens

import "sort"

func deleteStringIndex(s string, index int) string {
    return s[0:index] + s [index+1:]
}

func MinRemoveValidParens(arg string) string {
	openingStack := make([]int, 0)
	closingStack := make([]int, 0)

	for i, char := range arg {
		if char == ')' {
			stackLength := len(openingStack)
			// If the previous item added was NOT opening, then add the close for removal
			if stackLength == 0 {
				closingStack = append(closingStack, i)
			} else {
				// Pop off the latest opening
				openingStack = openingStack[:stackLength - 1]
			}
		} else if char == '(' {
			openingStack = append(openingStack, i)
		}
	}

	removals := append(openingStack, closingStack...)
	sort.Sort(sort.Reverse(sort.IntSlice(removals)))
	for _, index := range removals {
		arg = deleteStringIndex(arg, index)
	}

	return arg
}
