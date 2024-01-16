package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	str := "all is well"
	expectedMap := map[rune]int{'a': 1, 'l': 4, ' ': 2, 'i': 1, 's': 1, 'w': 1, 'e': 1}
	actualMap := parsing(str)
	assert.Equal(t, expectedMap, actualMap)
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
	expectedList := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	actualList := []int{}
	for th.Len() > 0 {
		t := heap.Pop(th)
		tree := t.(tree)
		actualList = append(actualList, tree.root.frequency)
	}

	assert.Equal(t, expectedList, actualList)
}
