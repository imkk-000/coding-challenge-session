package main_test

import (
	. "coding"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOddLinkedListInputLinkedList1ShouldBeLinkedListNil(t *testing.T) {
	inputData := &LinkedListNode{
		Data: 1,
		Next: nil,
	}

	actualLinkedList := RemoveOddLinkedList(inputData)

	assert.Nil(t, actualLinkedList)
}

func TestRemoveOddLinkedListInputLinkedList23ShouldBeLinkedList2(t *testing.T) {
	expectedLinkedList := []int{2}
	inputData := &LinkedListNode{
		Data: 2,
		Next: &LinkedListNode{
			Data: 3,
			Next: nil,
		},
	}

	list := RemoveOddLinkedList(inputData)
	var actualLinkedList []int
	for list != nil {
		actualLinkedList = append(actualLinkedList, list.Data)
		list = list.Next
	}

	assert.Equal(t, expectedLinkedList, actualLinkedList)
}

func TestRemoveOddLinkedListInputLinkedList1234ShouldBeLinkedList24(t *testing.T) {
	expectedLinkedList := []int{2, 4}
	inputData := &LinkedListNode{
		Data: 1,
		Next: &LinkedListNode{
			Data: 2,
			Next: &LinkedListNode{
				Data: 3,
				Next: &LinkedListNode{
					Data: 4,
					Next: nil,
				},
			},
		},
	}

	list := RemoveOddLinkedList(inputData)
	var actualLinkedList []int
	for list != nil {
		actualLinkedList = append(actualLinkedList, list.Data)
		list = list.Next
	}

	assert.Equal(t, expectedLinkedList, actualLinkedList)
}
