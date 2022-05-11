package min_remove_valid_parens_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/min_remove_valid_parens"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(MinRemoveValidParens("lee(t(c)o)de)")).To(Equal("lee(t(c)o)de"))
		Expect(MinRemoveValidParens("a)b(c)d")).To(Equal("ab(c)d"))
		Expect(MinRemoveValidParens("))((")).To(Equal(""))
		Expect(MinRemoveValidParens("")).To(Equal(""))
	})
})
