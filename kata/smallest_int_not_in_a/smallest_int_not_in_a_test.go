package smallest_int_not_in_a_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/smallest_int_not_in_a"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct index", func() {
		Expect(SmallestIntNotInA([]int{1, 3, 6, 4, 1, 2})).To(Equal(5))
		Expect(SmallestIntNotInA([]int{-1,-3})).To(Equal(1))
		Expect(SmallestIntNotInA([]int{1, 2, 3})).To(Equal(4))
	})
})
