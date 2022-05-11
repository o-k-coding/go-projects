package min_cost_climbing_stairs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/min_cost_climbing_stairs"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(MinCostClimbingStairs([]int{10,15,20})).To(Equal(15))
		Expect(MinCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1})).To(Equal(6))
	})
})
