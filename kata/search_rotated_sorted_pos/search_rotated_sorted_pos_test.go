package search_rotated_sorted_pos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/search_rotated_sorted_pos"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct index", func() {
		Expect(SearchRotatedSortedPos([]int{4,5,6,7,0,1,2}, 0)).To(Equal(4))
		Expect(SearchRotatedSortedPos([]int{4,5,6,7,0,1,2}, 3)).To(Equal(-1))
		Expect(SearchRotatedSortedPos([]int{1}, 0)).To(Equal(-1))
		Expect(SearchRotatedSortedPos([]int{1,3}, 3)).To(Equal(1))
	})
})
