package nth_tribonacci_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/nth_tribonacci"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(NthTribonacci(4)).To(Equal(4))
		Expect(NthTribonacci(25)).To(Equal(1389537))
	})
})
