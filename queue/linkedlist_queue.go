package queue

import "errors"

type Node struct {
	value int
	next  *Node
}

type LinkedListQueue struct {
	front    *Node
	back     *Node
	inserted int
}

func (queue *LinkedListQueue) Enqueue(value int) {
	newNode := &Node{value, nil}

	if queue.inserted == 0 {
		queue.front = newNode
		queue.back = newNode
	} else {
		queue.back.next = newNode
		queue.back = newNode
	}
	queue.inserted++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
	if queue.inserted == 0 {
		return -1, errors.New("Não se pode dar Dequeue em uma queue vazia")
	}

	value := queue.front.value

	if queue.inserted == 1 {
		queue.front = nil
		queue.back = nil
	} else {
		queue.front = queue.front.next
	}

	queue.inserted--
	return value, nil
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.inserted == 0 {
		return -1, errors.New("Não se pode dar Dequeue em uma queue vazia")
	}

	return queue.front.value, nil
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.inserted == 0
}

func (queue *LinkedListQueue) Size() int {
	return queue.inserted
}
