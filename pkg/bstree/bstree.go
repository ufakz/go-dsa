package bstree

//Node represents the component of a binary search tree.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert(Recursive) will add a node to the tree.
//The key to add should not already be in the tree.
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search
