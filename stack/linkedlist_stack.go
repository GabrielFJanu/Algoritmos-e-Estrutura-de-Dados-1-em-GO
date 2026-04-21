package stack

import (
	"errors"
)

type Node struct {
	value int
	next  *Node
}

type LinkedListStack struct {
	head     *Node
	inserted int
}

func (stack *LinkedListStack) Push(value int) {
	newNode := &Node{value, nil}

	newNode.next = stack.head
	stack.head = newNode // stack vazia e não fazia

	stack.inserted++
}

func (stack *LinkedListStack) Pop() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("Não se pode dar Pop em uma stack vazia.")
	}

	value := stack.head.value
	stack.head = stack.head.next
	stack.inserted--
	return value, nil
}

func (stack *LinkedListStack) Peek() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("Não se pode dar Pop em uma stack vazia.")
	}

	return stack.head.value, nil
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.inserted == 0
}

func (stack *LinkedListStack) Size() int {
	return stack.inserted
}
