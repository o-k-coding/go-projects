package valid_parentheses_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/valid_parentheses"
)

// Add test cases here

var _ = Describe("Basic tests ValidParenthesesRecursive", func() {
	It("Should return true for individual cases", func() {
		Expect(ValidParenthesesRecursive("()")).To(Equal(true))
		Expect(ValidParenthesesRecursive("[]")).To(Equal(true))
		Expect(ValidParenthesesRecursive("{}")).To(Equal(true))
	})

	It("Should return true for mixed case", func() {
		Expect(ValidParenthesesRecursive("()[]{}")).To(Equal(true))
		Expect(ValidParenthesesRecursive("(([]){})")).To(Equal(true))
	})

	It("Should return true for nested case", func() {
		Expect(ValidParenthesesRecursive("{[]}")).To(Equal(true))
	})

	It("Should return false for invalid case", func() {
		Expect(ValidParenthesesRecursive("(}")).To(Equal(false))
		Expect(ValidParenthesesRecursive("[]}")).To(Equal(false))
		Expect(ValidParenthesesRecursive("(]")).To(Equal(false))
		Expect(ValidParenthesesRecursive("()]")).To(Equal(false))
		Expect(ValidParenthesesRecursive("(}{)")).To(Equal(false))
	})
})

var _ = Describe("Basic tests ValidParenthesesWithStack", func() {
	It("Should return true for individual cases", func() {
		Expect(ValidParenthesesWithStack("()")).To(Equal(true))
		Expect(ValidParenthesesWithStack("[]")).To(Equal(true))
		Expect(ValidParenthesesWithStack("{}")).To(Equal(true))
	})

	It("Should return true for mixed case", func() {
		Expect(ValidParenthesesWithStack("()[]{}")).To(Equal(true))
		Expect(ValidParenthesesWithStack("(([]){})")).To(Equal(true))
	})

	It("Should return true for nested case", func() {
		Expect(ValidParenthesesWithStack("{[]}")).To(Equal(true))
	})

	It("Should return false for invalid case", func() {
		Expect(ValidParenthesesWithStack("(}")).To(Equal(false))
		Expect(ValidParenthesesWithStack("[]}")).To(Equal(false))
		Expect(ValidParenthesesWithStack("(]")).To(Equal(false))
		Expect(ValidParenthesesWithStack("()]")).To(Equal(false))
		Expect(ValidParenthesesWithStack("(}{)")).To(Equal(false))
	})
})
