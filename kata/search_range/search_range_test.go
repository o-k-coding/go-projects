package search_range_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/search_range"
)

// Add test cases here
var _ = Describe("Basic tests", func() {
	It("Should find the correct range", func() {
		Expect(SearchRange([]int{5,7,7,8,8,10}, 8)).To(Equal([]int{3,4}))
		Expect(SearchRange([]int{5,7,7,8,8,10}, 6)).To(Equal([]int{-1,-1}))
		Expect(SearchRange([]int{}, 0)).To(Equal([]int{-1,-1}))
	})
})
