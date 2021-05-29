package main

import (
	"github.com/rprajapati0067/datastructure-using-golang/linkedlist/singly"
)

func main() {
	singlylinkedLidt := singly.LinkedList{}

	singlylinkedLidt.InsertAtBegnning(1)
	singlylinkedLidt.InsertAtBegnning(2)
	singlylinkedLidt.InsertAtBegnning(3)
	singlylinkedLidt.InsertAtBegnning(4)
	singlylinkedLidt.InsertAtBegnning(5)

	singlylinkedLidt.Traverse()

}
