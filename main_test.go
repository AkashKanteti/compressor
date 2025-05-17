package main

import (
	"container/heap"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	str := "all is well"
	expectedMap := map[rune]int{'a': 1, 'l': 4, ' ': 2, 'i': 1, 's': 1, 'w': 1, 'e': 1}
	actualMap := parsing(str)
	assert.Equal(t, expectedMap, actualMap)
}

func TestParsingFromTxtFile(t *testing.T) {
	file, err := os.ReadFile("t.txt")
	assert.NoError(t, err, "error reading file")

	expectedMap := map[rune]int{'X': 333, 't': 223000}
	actualMap := parsing(string(file))
	assert.Equal(t, expectedMap['X'], actualMap['X'])
	assert.Equal(t, expectedMap['t'], actualMap['t'])
}

func TestTree(t *testing.T) {
	str := "all is well"
	mp := parsing(str)
	th := &treeHeap{}
	heap.Init(th)

	makeTree(mp, th)

	codes := map[rune]string{}

	rootany := heap.Pop(th)
	root := rootany.(tree)

	fmt.Printf("%v\n", string(root.root.rightnode.val))
	root.preOrder(root.root, codes, "")

	for k, v := range codes {
		fmt.Printf("%v : %v\n", k, v)
	}

}

func TestHeap(t *testing.T) {
	th := &treeHeap{}
	heap.Init(th)

	for i := 1; i <= 10; i++ {
		root1 := node{
			val:       1,
			frequency: i,
			isLeaf:    true,
		}
		tree1 := tree{
			root: &root1,
		}
		heap.Push(th, tree1)
	}
	expectedList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actualList := []int{}
	for th.Len() > 0 {
		t := heap.Pop(th)
		tree := t.(tree)
		actualList = append(actualList, tree.root.frequency)
	}

	assert.Equal(t, expectedList, actualList)
}

func TestPreOrder(t *testing.T) {
	th := &treeHeap{}
	heap.Init(th)

	for i := 1; i <= 10; i++ {
		root1 := node{
			val:       rune(i),
			frequency: i,
			isLeaf:    true,
			leftnode:  nil,
			rightnode: nil,
		}
		tree1 := tree{
			root: &root1,
		}
		heap.Push(th, tree1)
	}

	buildTree(th)

	codes := map[rune]string{}

	rootany := heap.Pop(th)
	root := rootany.(tree)

	root.preOrder(root.root, codes, "")

	for k, v := range codes {
		fmt.Printf("%v: %v\n", k, v)
	}

}
