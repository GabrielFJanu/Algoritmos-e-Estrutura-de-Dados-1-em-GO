package list

import "errors"

type Node2P struct {
	prev  *Node2P
	value int
	next  *Node2P
}

type DoublyLinkedList struct {
	head     *Node2P
	tail     *Node2P
	inserted int
}

func (list *DoublyLinkedList) Add(value int) {
	newNode := &Node2P{nil, value, nil}

	if list.inserted == 0 { // caso list vazia
		list.head = newNode
	} else { // caso list nao vazia
		list.tail.next = newNode
		newNode.prev = list.tail
	}

	list.tail = newNode

	list.inserted++
}

func (list *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo")
	}
	if index > list.inserted {
		return errors.New("index acima da faixa aceitável")
	}

	newNode := &Node2P{nil, value, nil}

	if list.inserted == 0 { // caso list vazia
		list.head = newNode
		list.tail = newNode
	} else {
		if index == 0 { // index no começo
			list.head.prev = newNode //independente
			newNode.next = list.head //independente
			list.head = newNode
		} else {
			prevNode := list.head
			for i := 0; i < index-1; i++ {
				prevNode = prevNode.next
			}
			if index == list.inserted {
				prevNode.next = newNode
				newNode.prev = prevNode
				list.tail = newNode
			} else {
				newNode.prev = prevNode
				newNode.next = prevNode.next
				prevNode.next = newNode
				newNode.next.prev = newNode
			}
		}
	}

	list.inserted++
	return nil
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo")
	}
	if index >= list.inserted {
		return errors.New("index acima da faixa aceitável")
	}

	if list.inserted == 1 {
		list.head = nil
		list.tail = nil
	} else {
		if index == 0 {
			list.head.next.prev = nil
			list.head = list.head.next
		} else {
			prevNode := list.head
			for i := 0; i < index-1; i++ {
				prevNode = prevNode.next
			}

			if index == list.inserted-1 {
				prevNode.next = nil
				list.tail = prevNode
			} else {
				prevNode.next = prevNode.next.next
				prevNode.next.prev = prevNode
			}
		}
	}

	list.inserted--
	return nil
}

func (list *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 {
		return -1, errors.New("index não pode ser negativo")
	}
	if index >= list.inserted {
		return -1, errors.New("index acima da faixa aceitável")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}
	return curNode.value, nil
}

func (list *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo")
	}
	if index >= list.inserted {
		return errors.New("index acima da faixa aceitável")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}
	curNode.value = value
	return nil
}

func (list *DoublyLinkedList) Size() int {
	return list.inserted
}

func (list *DoublyLinkedList) Reverse() {
	curNode := list.head

	for i := 0; i < list.inserted-1; i++ {
		curNode.next, curNode.prev = curNode.prev, curNode.next
		curNode = curNode.prev
	}

	list.head, list.tail = list.tail, list.head
}
