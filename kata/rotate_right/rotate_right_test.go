package rotate_right_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/rotate_right"
)


func createTestHead(nums []int) *ListNode {
	var head *ListNode
	// loop in reverse, for each one, create a new ListNode and set the Next equal to the reference to the previous value of head
	for i := (len(nums) - 1); i >= 0; i-- {
		num := nums[i]
		head = &ListNode{
			Val: num,
			Next: head,
		}
	}
	return head
}

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer", func() {
		head := createTestHead([]int{1,2,3,4,5,6})
		expected := createTestHead([]int{5,6,1,2,3,4})
		Expect(RotateRight(head, 2)).To(Equal(expected))
	})

	It("Should find the correct answer 2", func() {
		head := createTestHead([]int{0, 1, 2})
		expected := createTestHead([]int{2, 0, 1})
		Expect(RotateRight(head, 4)).To(Equal(expected))
	})

	It("Should find the correct answer 3", func() {
		head := createTestHead([]int{})
		expected := createTestHead([]int{})
		Expect(RotateRight(head, 0)).To(Equal(expected))
	})

	It("Should find the correct answer 4", func() {
		head := createTestHead([]int{1, 2})
		expected := createTestHead([]int{1, 2})
		Expect(RotateRight(head, 2)).To(Equal(expected))
	})

	It("Should find the correct answer 5", func() {
		head := createTestHead([]int{1, 2, 3, 4, 5})
		expected := createTestHead([]int{1, 2, 3, 4, 5})
		Expect(RotateRight(head, 10)).To(Equal(expected))
	})
})
