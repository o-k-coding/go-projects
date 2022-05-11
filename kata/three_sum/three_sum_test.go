package three_sum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/three_sum"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		// Expect(ThreeSum([]int{-1,0,1,2,-1,-4})).To(Equal([][]int{
		// 	{-1,0,1},
		// 	{-1,-1,2},
		// }))
		// Expect(ThreeSum([]int{})).To(Equal([][]int{}))
		// Expect(ThreeSum([]int{0})).To(Equal([][]int{}))
		// Expect(ThreeSum([]int{1, -1})).To(Equal([][]int{}))
		// Expect(ThreeSum([]int{0,0,0,0})).To(Equal([][]int{{0,0,0}}))

		Expect(ThreeSum([]int{3,0,-2,-1,1,2})).To(Equal([][]int{
			{-2,-1,3},
			{-2,0,2},
			{-1,0,1},
		}))
	})
})
