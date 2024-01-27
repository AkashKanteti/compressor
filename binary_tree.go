package main

type node struct {
	isLeaf    bool
	frequency int
	val       rune
	leftnode  *node
	rightnode *node
}

type tree struct {
	root *node
}

type treeHeap []tree

func (n treeHeap) Len() int {
	return len(n)
}

func (n treeHeap) Less(i, j int) bool {
	return n[i].root.frequency < n[j].root.frequency
}

func (n treeHeap) Swap(i, j int) {
	n[i].root, n[j].root = n[j].root, n[i].root
}

func (n *treeHeap) Push(x any) {
	*n = append(*n, x.(tree))
}

func (n *treeHeap) Pop() any {
	len := n.Len()
	old := *n
	x := old[len-1]
	*n = old[0 : len-1]
	return x
}

func (t tree) preOrder(root *node, m map[rune]string, val string) {
	if root == nil {
		return
	}
	leftstring := val
	rightstring := val
	t.preOrder(root.leftnode, m, leftstring+"0")
	t.preOrder(root.rightnode, m, rightstring+"1")
	if root.isLeaf {
		m[root.val] = val
	}
}
