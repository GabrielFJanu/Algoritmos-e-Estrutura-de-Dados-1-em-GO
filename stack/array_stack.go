package stack

import "errors"

type ArrayStack struct {
	values   []int
	inserted int
}

func (stack *ArrayStack) Init(size int) error {
	if size <= 0 {
		return errors.New("Size não pode ser um valor <= 0.")
	}

	stack.values = make([]int, size)
	return nil
}

func (stack *ArrayStack) doubleArray() {
	newArray := make([]int, 2*len(stack.values))

	for i := 0; i < len(stack.values); i++ {
		newArray[i] = stack.values[i]
	}

	stack.values = newArray
}

func (stack *ArrayStack) Push(value int) {
	if stack.inserted == len(stack.values) {
		stack.doubleArray()
	}

	stack.values[stack.inserted] = value // lista vazia e não vazia

	stack.inserted++
}

func (stack *ArrayStack) Pop() (int, error) {
	if stack.inserted == 0 { // lista vazia
		return -1, errors.New("Não se pode dar Pop em stack vazia.")
	}

	stack.inserted--
	return stack.values[stack.inserted], nil
}

func (stack *ArrayStack) Peek() (int, error) {
	if stack.inserted == 0 { // lista vazia
		return -1, errors.New("Não se pode dar Pop em stack vazia.")
	}

	return stack.values[stack.inserted-1], nil
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.inserted == 0
}

func (stack *ArrayStack) Size() int {
	return stack.inserted
}
