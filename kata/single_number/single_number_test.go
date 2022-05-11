package single_number_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/single_number"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(SingleNumber([]int{2,2,1})).To(Equal(1))
		Expect(SingleNumber([]int{4,1,2,1,2})).To(Equal(4))
		Expect(SingleNumber([]int{1})).To(Equal(1))
	})
})
