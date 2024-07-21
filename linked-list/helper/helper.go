package helper

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}
type Head struct {
	Head *LinkedList
}

func (h *Head) InsertValueInList() {
	slice := []int{1, 4, 8, 2, 5, 9, 3, 6}
	currentNode := h.Head
	for i := 0; i < len(slice); i++ {
		node := &LinkedList{}
		node.Value = slice[i]
		node.Next = nil
		if currentNode == nil {
			h.Head = node
			currentNode = node
		} else {
			currentNode.Next = node
			currentNode = node
		}
	}
	fmt.Println("Inserted Node and return head ", h.Head)
}

func (h *Head) GetAllNode() []int {
	currentNode := h.Head
	var slices []int
	for currentNode != nil {
		slices = append(slices, currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("Get all nodes ", slices)
	return slices
}

func (h *Head) GetNumberOfNodeInList() {
	cusrrentNode := h.Head
	count := 0
	for cusrrentNode != nil {
		cusrrentNode = cusrrentNode.Next
		count++
	}
	fmt.Println("Number of Node in List ", count)
}

func (h *Head) DeleteFirstNode() {
	currentNode := h.Head
	h.Head = currentNode.Next
	currentNode = h.Head
	var slice []int
	for currentNode != nil {
		slice = append(slice, currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("Delete First Node from list ", slice)
}

func (h *Head) DeleteLastNode() {
	nextNode := h.Head
	preNode := nextNode
	for nextNode.Next != nil {
		preNode = nextNode
		nextNode = nextNode.Next
	}
	preNode.Next = nil
	var slice []int
	currentNode := h.Head
	for currentNode != nil {
		slice = append(slice, currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("Delete Last Node from List ", slice)
}

func (h *Head) DeleteFromIndex(idx int) {
	if idx < 0 {
		fmt.Println("Invalid Index")
	} else if idx == 0 {
		h.Head = h.Head.Next
	}
	if idx > 0 {
		currentNode := h.Head
		preNode := currentNode
		count := 0
		for currentNode != nil {
			if count == idx {
				preNode.Next = currentNode.Next
				break
			}
			preNode = currentNode
			currentNode = currentNode.Next
			count++
		}
	}
	currentNode := h.Head
	var slice []int
	for currentNode != nil {
		slice = append(slice, currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Printf("Delete from indext %d, result = %v", idx, slice)
}
