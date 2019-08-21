package main

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

func RemoveOddLinkedList(headList *LinkedListNode) *LinkedListNode {
	var previousNode *LinkedListNode
	currentNode := headList
	for currentNode != nil {
		if currentNode.Data%2 == 1 {
			if previousNode == nil {
				headList = headList.Next
				currentNode = headList
				continue
			}
			previousNode.Next = currentNode.Next
			currentNode = currentNode.Next
			continue
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return headList
}
