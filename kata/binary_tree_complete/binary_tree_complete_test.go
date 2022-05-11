package binary_tree_complete_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/binary_tree_complete"
)

func populateNode(nums []int, i int) *TreeNode {
	// This problem specifically works with numbers > 0, so 0 is our "null" leaf placeholder
	if nums[i] == 0 {
		return nil
	}
	length := len(nums)
	node := &TreeNode{
		Val: nums[i],
	}
	leftI := 2 * i + 1
	rightI := 2 * i + 2
	if leftI < length {
		node.Left = populateNode(nums, leftI)
	}
	if rightI < length {
		node.Right = populateNode(nums, rightI)
	}
	return node
}

var _ = Describe("Basic tests", func() {
	It("Should return true for a complete tree", func() {
		testTree := populateNode([]int{1,2,3,4,5,6}, 0)
		Expect(BinarySearchTreeComplete(testTree)).To(Equal(true))
	})

	It("Should return false for an incomplete tree", func() {
		testTree := populateNode([]int{1,2,3,4,5,0,7}, 0)
		Expect(BinarySearchTreeComplete(testTree)).To(Equal(false))
	})
})
