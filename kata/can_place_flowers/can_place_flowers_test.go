package can_place_flowers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/can_place_flowers"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(CanPlaceFlowers([]int{1,0,0,0,1}, 1)).To(Equal(true))
		Expect(CanPlaceFlowers([]int{1,0,0,0,1}, 2)).To(Equal(false))
		Expect(CanPlaceFlowers([]int{1,0,0,0,0,1}, 2)).To(Equal(false))
	})
})
