package valid_parentheses_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestValidParentheses(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "valid_parentheses Suite")
}
