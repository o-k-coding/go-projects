package remove_dupes_sorted_ii_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/remove_dupes_sorted_ii"
)

func printListNode(head *ListNode) {
	fmt.Println("Printing linked list")
	for head != nil {
		fmt.Printf("val %d \n", head.Val)
		head = head.Next
	}
	fmt.Println("Done Printing linked list")
}

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

	printListNode(head)

	return head
}


// Add test cases here

var _ = Describe("Basic tests", func() {
	It("Should find the correct answer 1", func() {
		testHead := createTestHead([]int{1,2,3,3,4,4,5})
		expectedHead := createTestHead([]int{1, 2, 5})
		Expect(RemoveDupesSortedIi(testHead)).To(Equal(expectedHead))
	})

	It("Should find the correct answer 2", func() {
		testHead := createTestHead([]int{1,1,1,2,3})
		expectedHead := createTestHead([]int{2, 3})
		Expect(RemoveDupesSortedIi(testHead)).To(Equal(expectedHead))
	})
})
