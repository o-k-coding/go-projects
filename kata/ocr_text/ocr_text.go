package ocr_text

import (
	"unicode"
)


func ExpandUnknownChars(text string) string {
	expandedText := ""
	for _, char := range text {
		isNumber := unicode.IsNumber(char)
		if isNumber {
			// Add a 1 character n times, where n is the number represented by the numeric character
			unknownNum := int(char - '0')
			expandedUnknowns := ""
			for len(expandedUnknowns) < unknownNum {
				expandedUnknowns += "1"
			}
			expandedText += expandedUnknowns
		} else {
			expandedText += string(char)
		}
	}
	return expandedText
}

func OcrText(S string, T string) bool {
	if S == T {
		return true
	}

	// First expand the unknowns into the correct number of consistent characters for both strings
	sResults := ExpandUnknownChars(S)
	tResults := ExpandUnknownChars(T)

	// At this point, if the lengths do not match, then there is no way the strings could match
	if len(sResults) != len(tResults) {
		return false
	}


	// Since both strings are the same length we can use 1 index to loop over them
	for i := 0; i < len(sResults); i ++ {
		sChar := sResults[i]
		tChar := tResults[i]
		// Ff either char is unknown then we could still have a valid match
		if (sChar == '1' || tChar == '1') {
			continue
		}
		// If both characters are known, and they are not equal, then the strings cannot match, we can short circuit here
		if (sChar != tChar) {
			return false
		}
	}

	// If we get to this point, then the strings could be a valid match
	return true
}
