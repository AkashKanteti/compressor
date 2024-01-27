package main

import (
	"container/heap"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	mp := parsing(string(file))

	fmt.Printf("%v", mp)

	th := &treeHeap{}
	heap.Init(th)

	for k, v := range mp {
		fmt.Printf("%v : %v\n", string(k), v)
		node := node{
			frequency: v,
			val:       k,
			isLeaf:    true,
		}

		tree := tree{
			root: &node,
		}

		heap.Push(th, tree)
	}

	for th.Len() > 1 {
		t1 := heap.Pop(th)
		t2 := heap.Pop(th)

		tree1 := t1.(tree)
		tree2 := t2.(tree)

		node := node{
			frequency: tree1.root.frequency + tree2.root.frequency,
			isLeaf:    false,
		}

		treef := tree{
			root:      &node,
			leftnode:  tree1.root,
			rightnode: tree2.root,
		}

		heap.Push(th, treef)
	}

}

func parsing(text string) map[rune]int {

	mp := make(map[rune]int)
	for _, ch := range text {
		mp[ch]++
	}
	return mp
}
