package validate

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString // any lowercase letter, number or underscore 0 or more times
	// Note this is not a complete solution
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s'-]+$`).MatchString // any lowercase letter, number or underscore 0 or more times
)

func ValdidateString(val string, minLen int, maxLen int) error {
	n := len(val)
	if n < minLen || n > maxLen {
		return fmt.Errorf("must contain from %d-%d characters", minLen, maxLen)
	}
	return nil
}

func ValidateUsername(val string) error {
	if err := ValdidateString(val, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(val) {
		return fmt.Errorf("must contain only letters, spaces, apostrophe (') or dash (-)")
	}
	return nil
}

func ValidatePassword(val string) error {
	return ValdidateString(val, 6, 100)
}

func ValidateEmail(val string) error {
	if err := ValdidateString(val, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(val); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}

func ValidateFullName(val string) error {
	if err := ValdidateString(val, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(val) {
		return fmt.Errorf("must contain only lowercase letters, digits or underscore")
	}
	return nil
}
