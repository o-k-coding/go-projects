package cherry_pickup_ii_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/cherry_pickup_ii"
)

// Add test cases here

var _ = Describe("Basic tests", func() {

	It("Should find the max cherries simple", func() {
		Expect(CherryPickup([][]int{
			{3,1,1},
		})).To(Equal(4))


		Expect(CherryPickup([][]int{
			{3,1,1},
			{2,5,1},
		})).To(Equal(4 + 7))

		Expect(CherryPickup([][]int{
			{3,1,1},
			{2,5,1},
			{1,5,5},
		})).To(Equal(4 + 7 + 10))

		Expect(CherryPickup([][]int{
			{3,1,1},
			{2,5,1},
			{1,5,5},
			{2,1,1},
		})).To(Equal(4 + 7 + 10 + 3))
	})

	It("Should find the max cherries complex", func() {
		grid := [][]int{
			{1,0,0,0,0,0,1}, // 2
			{2,0,0,0,0,3,0}, // 7
			{2,0,9,0,0,0,0}, // 9
			{0,3,0,5,4,0,0}, // 16
			{1,0,2,3,0,0,6},
		}

		Expect(CherryPickup(grid)).To(Equal(28))
	})
})
