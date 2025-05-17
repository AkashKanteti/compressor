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

	for k, v := range codes {
		fmt.Printf("%v : %v\n", string(k), v)
	}

	//converting codes to codes: TODO do this when preorder
	newCodes := covertToBytes(codes)

	fmt.Print("new code\n")
	for k, v := range newCodes {
		fmt.Printf("%v : %v\n", string(k), v)
	}

	writeToFile(newCodes, string(file))
	fmt.Printf("done")
}

func covertToBytes(codes map[rune]string) map[rune][]byte {
	newCodes := map[rune][]byte{}
	for k, code := range codes {
		sikz := len(code)/8 + min(len(code)%8, 1)
		byt := make([]byte, sikz)

		now := 0
		ok := uint8(0)
		idx := 0
		fmt.Printf("%v\n", code)
		for i := len(code) - 1; i >= 0; i -= 1 {

			if code[i] == '1' {
				ok += 1 << now
			}
			now++

			fmt.Printf("%v\n", ok)

			if now == 8 {
				byt[idx] = ok
				idx++
				now = 0
				ok = 0
			}
		}

		if now != 1 && ok > 0 {
			// fmt.Printf("%v\n", now)
			byt[idx] = ok
		}
		newCodes[k] = byt

	}

	// fmt.Printf("%v and %v\n", byt, string(byt))
	// for _, v := range string(byt) {
	// 	fmt.Printf("%v\n", rune(v))
	// }
	// fmt.Printf("%v", len(byt))

	return newCodes
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

	buildTree(th)
}

func buildTree(th *treeHeap) {
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
func writeToFile(codes map[rune][]byte, oldfile string) {
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
		str += string(k) + "," + string(v) + " "
	}

	//payload
	payload := []byte{}
	for _, ch := range oldfile {
		fmt.Printf("%v: %v\n", ch, codes[ch])
		payload = append(payload, codes[ch]...)
	}

	_, err = file.Write(payload)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("done")
}
