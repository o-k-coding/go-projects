package rotate_right

type ListNode struct {
	Val int;
	Next *ListNode;
}

func RotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	nodeByPosition := make(map[int]*ListNode)
	tracker := head
	length := 0

	// First find the length and store each node by index for easier retrieval
	for tracker != nil {
		nodeByPosition[length] = tracker
		tracker = tracker.Next
		length ++
	}

	// Now adjust k for the length of the list if needed
	if k >= length {
		k = k % length
	}

	if k == 0 {
		return head
	}

	// Next get the new head
	newHeadIndex := length - k
	newHead := nodeByPosition[newHeadIndex]

	// Next hand the old head off to the previous tail
	oldTail := nodeByPosition[length - 1]
	oldTail.Next = head

	// Next clear the new tail
	newTail := nodeByPosition[newHeadIndex - 1]
	newTail.Next = nil
	return newHead
}
