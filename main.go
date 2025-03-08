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

	// fmt.Printf("%v", mp)

	th := &treeHeap{}
	heap.Init(th)

	makeTree(mp, th)

	codes := map[rune]string{}

	rootany := heap.Pop(th)
	root := rootany.(tree)

	root.preOrder(root.root, codes, "")

	// for k,v:=range codes{
	// 	fmt.Printf("%v : %v\n",string(k),v)
	// }

	writeToFile(codes, string(file))
	fmt.Printf("done")
}

func parsing(text string) map[rune]int {

	mp := make(map[rune]int)
	for _, ch := range text {
		mp[ch]++
	}
	return mp
}

func makeTree(mp map[rune]int, th *treeHeap) {
	for k, v := range mp {
		node := node{
			frequency: v,
			val:       k,
			isLeaf:    true,
			leftnode:  nil,
			rightnode: nil,
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
			leftnode:  tree1.root,
			rightnode: tree2.root,
		}

		treef := tree{
			root: &node,
		}

		heap.Push(th, treef)
	}
}

func writeToFile(codes map[rune]string, oldfile string) {
	file, err := os.Create("C:/Users/LEGION/Desktop/compressor/output.txt")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	defer file.Close()

	//prepare header

	str := ""
	str = fmt.Sprintf("%v", len(codes)) + " "
	for k, v := range codes {
		str += string(k) + "," + v + " "
	}

	//payload
	for _, ch := range oldfile {
		fmt.Printf(codes[ch] + "\n")
		str += codes[ch]
	}

	_, err = file.Write([]byte(str))
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	// fmt.Printf("done")
}
