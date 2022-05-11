package remove_dupes_sorted_ii

type ListNode struct {
	Val int
	Next *ListNode
}

// Delete all nodes that have duplicate numbers, leaving only distint numbers
// Result must be sorted
// Input is sorted
func RemoveDupesSortedIi(head *ListNode) *ListNode {
	var result *ListNode = &ListNode{
		Next: head,
	}
	var tracker *ListNode = result
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			// For each node where the head equals the next, skip
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// Once we have skipped all of the duplicates, "hand off" the good node to the tracker to "delete" the duplicates
			tracker.Next = head.Next
		} else {
			tracker = tracker.Next
		}
		head = head.Next
	}
	return result.Next
}
