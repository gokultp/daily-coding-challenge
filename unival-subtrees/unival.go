package unival

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) CountUnivalSubtrees() int {
	if t == nil || t.Root == nil {
		return 0
	}
	_,_, count := t.Root.IsUnival()
	return count
}

func (n *Node) IsUnival()(bool, int, int) {
	if n.Left == nil && n.Right == nil {
		return true, n.Value, 1
	}
	if n.Left == nil {
		_, _,rCount := n.Right.IsUnival()
		return false, -999999, rCount
	}
	if n.Right == nil {
		_, _, lCount := n.Left.IsUnival()
		return false, -999999, lCount 
	}
	rUnival, rVal,rCount := n.Right.IsUnival()
	lUnival, lVal,lCount := n.Left.IsUnival()
	if lUnival && rUnival && lVal == n.Value && rVal == n.Value{
		return true, n.Value, rCount + lCount +1

	}
	return false, -999999, rCount + lCount

}