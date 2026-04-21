package list

import (
	"errors"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head     *Node
	inserted int
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{value, nil}

	curNode := list.head
	for i := 0; i < list.inserted; i++ {
		curNode = curNode.next
	}
	curNode.next = newNode

	list.inserted++
}

func (list *LinkedList) AddOnIndex(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo.")
	}
	if index > list.inserted {
		return errors.New("index acima do range aceitável.")
	}

	newNode := &Node{value, nil}
	if index == 0 { // também cobre o caso que a lista está vazia
		newNode.next = list.head
		list.head = newNode
	} else {
		prevNode := list.head
		for i := 0; i < index-1; i++ {
			prevNode = prevNode.next
		}

		if index == list.inserted { //fim
			prevNode.next = newNode
		} else {
			newNode.next = prevNode.next
			prevNode.next = newNode
		}
	}
	list.inserted++
	return nil
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo.")
	}
	if index >= list.inserted {
		return errors.New("index acima do range aceitável.")
	}

	if index == 0 { // começo e lista com um único elemento
		list.head = list.head.next // também cobre o caso do único elemento
	} else { // meio e fim
		prevNode := list.head
		for i := 0; i < index-1; i++ {
			prevNode = prevNode.next
		}
		prevNode.next = prevNode.next.next // também combre o caso do fim
	}

	list.inserted--
	return nil
}

func (list *LinkedList) Get(index int) (int, error) {
	if index < 0 {
		return -1, errors.New("index não pode ser negativo.")
	}
	if index >= list.inserted {
		return -1, errors.New("index acima do range aceitável.")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}
	return curNode.value, nil
}

func (list *LinkedList) Set(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo.")
	}
	if index >= list.inserted {
		return errors.New("index acima do range aceitável.")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}
	curNode.value = value
	return nil
}

func (list *LinkedList) Size() int {
	return list.inserted
}

func (list *LinkedList) Reverse() {
	var prevNode *Node = nil
	curNode := list.head

	for curNode != nil {
		next := curNode.next
		curNode.next = prevNode
		prevNode = curNode
		curNode = next
	}

	list.head = prevNode
}
