package main

import "fmt"

// 有序双向链表
type LinkNode struct {
	Score      int
	Value      string
	Next, Prev *LinkNode
}

func (n *LinkNode) Add(score int, value string) {
	// 插入
	newNode := &LinkNode{
		Score: score,
		Value: value,
	}
	if n == nil {
		n = newNode
		return
	}
	// find insert Idx
	insertBeforeNode := n
	for insertBeforeNode.Score < score {
		if insertBeforeNode.Next == nil {
			insertBeforeNode.Next = newNode
			newNode.Prev = insertBeforeNode
			return
		}
		insertBeforeNode = insertBeforeNode.Next
	}

	newNode.Next = insertBeforeNode
	insertBeforeNode.Prev.Next = newNode
	return
}

func (n *LinkNode) Print() {
	curNode := n
	for curNode != nil {
		fmt.Println("socre:%d, value:%s", curNode.Score, curNode.Value)
		curNode = curNode.Next
	}
}

func testSortLink() {
	head := new(LinkNode)
	head.Add(97, "gaoda")
	head.Add(100, "chenhua")
	head.Add(1, "liudehua")
	head.Print()
}
