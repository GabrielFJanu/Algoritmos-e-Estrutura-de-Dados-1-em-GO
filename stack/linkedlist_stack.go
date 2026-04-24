package stack

import "errors"

type Node struct {
	value int
	next  *Node
}

type LinkedListStack struct {
	top      *Node
	inserted int
}

func (stack *LinkedListStack) Push(value int) {
	newNode := &Node{value, nil}
	newNode.next = stack.top
	stack.top = newNode

	stack.inserted++
}

func (stack *LinkedListStack) Pop() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("Não se pode dar Pop em uma stack vazia")
	}

	value := stack.top.value

	stack.top = stack.top.next

	stack.inserted--
	return value, nil
}

func (stack *LinkedListStack) Peek() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("Não se pode dar Pop em uma stack vazia")
	}

	return stack.top.value, nil
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.inserted == 0
}

func (stack *LinkedListStack) Size() int {
	return stack.inserted
}
