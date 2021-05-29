package singly

import "fmt"

type ListNode struct {
	data interface{}
	next *ListNode
}
type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) Traverse() {

	if ll.head == nil {
		fmt.Println("traversing : List is empty")
		return
	}
	current := ll.head
	for current != nil {

		if ll.head == nil {

		} else {
			fmt.Printf("%v -> ", current.data)
		}

		current = current.next
	}
	return
}

func (ll *LinkedList) InsertAtBegnning(data interface{}) {
	node := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
	return
}
