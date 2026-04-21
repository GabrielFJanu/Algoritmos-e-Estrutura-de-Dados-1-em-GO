package list

import "errors"

type ArrayList struct {
	values   []int
	inserted int
}

func (list *ArrayList) Init(size int) error {
	if size <= 0 {
		return errors.New("size não pode ser <= 0.")
	}
	list.values = make([]int, size)
	return nil
}

func (list *ArrayList) doubleArray() {
	newArray := make([]int, 2*len(list.values))
	for i := 0; i < len(list.values); i++ {
		newArray[i] = list.values[i]
	}
	list.values = newArray
}

func (list *ArrayList) Add(value int) {
	if list.inserted == len(list.values) {
		list.doubleArray()
	}

	list.values[list.inserted] = value

	list.inserted++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo.")
	}
	if index > list.inserted {
		return errors.New("index acima do range aceitável.")
	}

	if list.inserted == len(list.values) {
		list.doubleArray()
	}

	for i := list.inserted; i > index; i-- {
		list.values[i] = list.values[i-1]
	}
	list.values[index] = value

	list.inserted++
	return nil
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo.")
	}
	if index >= list.inserted {
		return errors.New("index acima do range aceitável.")
	}

	for i := index; i < list.inserted-1; i++ {
		list.values[i] = list.values[i+1]
	}

	list.inserted--
	return nil
}

func (list *ArrayList) Get(index int) (int, error) {
	if index < 0 {
		return -1, errors.New("index não pode ser negativo.")
	}
	if index >= list.inserted {
		return -1, errors.New("index acima do range aceitável.")
	}

	return list.values[index], nil
}

func (list *ArrayList) Set(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo.")
	}
	if index >= list.inserted {
		return errors.New("index acima do range aceitável.")
	}

	list.values[index] = value
	return nil
}

func (list *ArrayList) Size() int {
	return list.inserted
}
