package binary_search_tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/binary_search_tree"
)


// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		Expect(BinarySearchTree("")).To(Equal(""))
	})
})
