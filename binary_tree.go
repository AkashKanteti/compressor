package main

type node struct {
	isLeaf bool
	val    int
}

type tree struct {
	root      node
	leftnode  node
	rightnode node
}

type treeHeap []tree

func (n *treeHeap) Len() int {
	return len(*n)
}

func (n *treeHeap) Less(i, j tree) bool {
	if i.root.val > j.root.val {
		return true
	}
	return false
}

func (n *treeHeap) Swap(i, j tree) {
	i.root.val, j.root.val = j.root.val, i.root.val
	i.root.isLeaf, j.root.isLeaf = j.root.isLeaf, i.root.isLeaf
}

func (n *treeHeap) Push(x tree) {
	*n = append(*n, x)
}

func (n *treeHeap) Pop() any {
	len := n.Len()
	old := *n
	x := old[len-1]
	*n = old[0 : len-1]
	return x
}
