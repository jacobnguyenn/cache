package lrucache

import (
	"fmt"
)

type DoublyLinkedList struct {
	len  int
	tail *ListNode
	head *ListNode
}

type ListNode struct {
	Val  any
	Prev *ListNode
	Next *ListNode
}

func NewDoubleLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}
func (c *DoublyLinkedList) PushBack(node *ListNode) {
	// empty list
	if c.head == nil {
		c.head = node
		c.tail = node
		c.len++
		return
	}
	c.tail.Next = node
	node.Prev = c.tail
	node.Next = nil
	c.tail = node
	c.len++
}

func (c *DoublyLinkedList) AddToHead(node *ListNode) {
	if c.head == nil {
		c.head = node
		c.tail = node
		return
	}
	node.Next = c.head
	c.head.Prev = node
	c.head = node
}

func (c *DoublyLinkedList) PopHead() {
	// empty list
	if c.head == nil {
		return
	}
	// 1 element list
	if c.head.Next == nil {
		c.head = nil
		c.len--
		return
	}
	c.head = c.head.Next
	c.head.Next.Prev = nil
	c.len--
}

func (c *DoublyLinkedList) Size() int {
	if c.head == nil {
		return 0
	}
	var size int
	var cur *ListNode
	for cur = c.head; cur != nil; cur = cur.Next {
		size++
	}
	return size
}

func (c *DoublyLinkedList) MoveToEnd(node *ListNode) {
	var cur *ListNode
	if c.head == nil {
		c.head = node
		c.tail = node
		return
	}
	if c.head == node {
		return
	}
	for cur = c.head; cur != nil; cur = cur.Next {
		if cur == node {
			// first remove this element
			cur.Prev.Next = cur.Next
			cur.Next.Prev = cur.Prev
			// then move it to the end
			c.PushBack((node))
		}
	}
}

func (c *DoublyLinkedList) Print() {
	if c.head == nil {
		fmt.Print("")
	}
	var count int
	var cur *ListNode
	for cur = c.head; cur != nil; cur = cur.Next {
		count++
		if count == c.Size() {
			fmt.Printf("val %v\n", cur.Val)
			return
		}
		fmt.Printf("val %v -> ", cur.Val)
	}
}
