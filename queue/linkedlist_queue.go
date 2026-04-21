package queue

import "errors"

type Node struct {
	value int
	next  *Node
}

type LinkedListQueue struct {
	head     *Node
	tail     *Node
	inserted int
}

func (queue *LinkedListQueue) Enqueue(value int) {
	newNode := &Node{value, nil}

	if queue.inserted == 0 { // queue vazia
		queue.head = newNode
	} else { // queue não vazia
		queue.tail.next = newNode
	}
	queue.tail = newNode

	queue.inserted++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
	if queue.inserted == 0 {
		return -1, errors.New("Não se pode dar Dequeue em uma queue vazia.")
	}

	value := queue.head.value

	if queue.inserted == 1 { // elemento único
		queue.tail = nil
	}
	queue.head = queue.head.next // elemento único e resto

	queue.inserted--
	return value, nil
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.inserted == 0 {
		return -1, errors.New("Não se pode dar Front em uma queue vazia.")
	}

	return queue.head.value, nil
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.inserted == 0
}

func (queue *LinkedListQueue) Size() int {
	return queue.inserted
}
