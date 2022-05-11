package valid_parentheses

var closingsRec = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

var closings = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func ValidParenthesesRecursive(s string) bool {
	// Go until you find a closer with an opening right before it
	// if you find it, remove that one and re call this function with the string minus those two chars
	// return false if you can't find one at all

	// Base cases
	if len(s) == 0 {
		return true
	} else if len(s) % 2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		// if you find a closing paren
		if open, ok := closingsRec[s[i]]; ok {

			// if the closer is the start, we can't have a valid string
			if i == 0 {
				return false
			}
			// If the previous character is the matching opening
			if s[i -1] == open {
				// cut out the matching parens and check the remaining string
				sub := s[: i - 1] + s[i + 1:]
				return ValidParenthesesRecursive(sub)
			}
		}
	}

	return false
}

func ValidParenthesesWithStack(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) % 2 != 0 {
		return false
	}

	matchStack := make([]rune, 0)

	for _, char := range s {
		if match, ok := closings[char]; ok {
			stackLen := len(matchStack)
			// If a closing was found first, then we can't have a valid string
			if stackLen == 0 {
				return false
			}

			// check if the previous item in the stack was the correct opening
			if matchStack[stackLen - 1] == match {
				// pop this one off the stack
				matchStack = matchStack[:stackLen - 1]
			} else {
				// If the previous item was not a match, then we can't have a valid string
				return false
			}
		} else {
			// Add one to the stack for cheecking
			matchStack = append(matchStack, char)
		}
	}

	return len(matchStack) == 0
}
