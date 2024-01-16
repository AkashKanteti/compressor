package main

type node struct {
	isLeaf    bool
	frequency int
	val       rune
}

type tree struct {
	root      *node
	leftnode  *node
	rightnode *node
}

type treeHeap []tree

func (n treeHeap) Len() int {
	return len(n)
}

func (n treeHeap) Less(i, j int) bool {
	return n[i].root.frequency > n[j].root.frequency
}

func (n treeHeap) Swap(i, j int) {
	n[i].root, n[j].root = n[j].root, n[i].root
	n[i].leftnode, n[j].leftnode = n[j].leftnode, n[i].leftnode
	n[i].rightnode, n[j].rightnode = n[j].rightnode, n[i].rightnode
	// n[i].root.frequency, n[j].root.frequency = n[j].root.frequency, n[i].root.frequency
	// n[i].root.isLeaf, n[j].root.isLeaf = n[j].root.isLeaf, n[i].root.isLeaf
	// n[i].root.val, n[j].root.val = n[j].root.val, n[i].root.val
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
