package search_2d_matrix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/search_2d_matrix"
)

// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(Search2dMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3)).To(Equal(true))
		Expect(Search2dMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13)).To(Equal(false))
	})
})
