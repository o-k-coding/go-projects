package majority_element_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/majority_element"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(MajorityElement([]int{3,2,3})).To(Equal(3))
		Expect(MajorityElement([]int{2,2,1,1,1,2,2})).To(Equal(2))
	})
})

// c 0
// 3 is new candidate
// c 1 incremement because num is candidate

// c 1
// no new candidate
// c 0 decrement because num is not candidate

// c 0
// 3 is new candidate
// c 1 increment because num is candidate
