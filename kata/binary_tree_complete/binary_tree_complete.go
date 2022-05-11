package binary_tree_complete

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func BinarySearchTreeComplete(root *TreeNode) bool {

	queue := []*TreeNode{root}
	emptyNodeFound := false
	// set to 2 to allow the root to pass
	prevChildrenCount := 2

	// For each node, queue up the two children
	// If at any time
	for len(queue) > 0 {
		node := queue[0]

		// If the left node is missing but the right isnt't, game over
		if node.Left == nil && node.Right != nil {
			return false
		}

		childrenCount := 0

		if node.Left != nil {
			// If we have previously seen an empty node on this level, we are done
			if emptyNodeFound {
				// Check here to allow the empty node found to be set once and only once
				// Since we are doing BFS, the first time we see an empty node should be at the very bottom, and only once, by then we should be done
				//
				return false
			}
			queue = append(queue, node.Left)
			childrenCount ++
		} else {
			emptyNodeFound = true
		}
		if node.Right != nil {
			// If we have previously seen an empty node, we are done, this is because for a tree to be complete, the empty should be the "final" leaf
			// which will show up once, then we should be able to complete the queue after that.
			if emptyNodeFound {
				return false
			}
			queue = append(queue, node.Right)
			childrenCount ++
		} else {
			emptyNodeFound = true
		}

		// If the child count is greater than the previous node, that means there is an empty node to the left or above this node
		if childrenCount > prevChildrenCount {
			return false
		}
		// Remove the visited node
		queue = queue[1:]
	}

	return true
}

// 0
// 1 2
// n 3 4 5

// 0 1 - needs 1
// 1 2
// 3 n 4 5
