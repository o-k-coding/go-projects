package maximize_person_distance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/maximize_person_distance"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(MaximizePersonDistance([]int{0,1,0,0,0,0})).To(Equal(4))
		Expect(MaximizePersonDistance([]int{1,0,0,0,1,0,1})).To(Equal(2))
		Expect(MaximizePersonDistance([]int{1,0,0,0,0,1,0,1})).To(Equal(2))
		Expect(MaximizePersonDistance([]int{1,0,0,0})).To(Equal(3))
		Expect(MaximizePersonDistance([]int{0,1})).To(Equal(1))
	})
})
