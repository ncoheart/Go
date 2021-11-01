package main

import "fmt"

// 定义一个单向链表
type Node struct {
	data int
	next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2

	head.next = node1
	var node2 = new(Node)
	node2.data = 3

	node1.next = node2
	ShowNode(head)

	// 头部插入新节点
	var tail *Node
	tail = head
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		node.next = tail
		tail = &node
	}
	ShowNode(tail)
}
