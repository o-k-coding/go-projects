package search_insert_pos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/search_insert_pos"
)

// Add test cases here
var _ = Describe("Basic tests", func() {
	It("Should find the correct index", func() {
		Expect(SearchInsertPos([]int{1,3,5,6}, 5)).To(Equal(2))
		Expect(SearchInsertPos([]int{1,3,5,6}, 2)).To(Equal(1))
		Expect(SearchInsertPos([]int{1,3,5,6}, 7)).To(Equal(4))
	})
})
