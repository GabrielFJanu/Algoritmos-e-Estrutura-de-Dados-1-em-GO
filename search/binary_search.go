package search

func IterativeBinarySearch(array []int, value int) int {
	ini := 0
	fim := len(array) - 1

	for ini <= fim {
		meio := (ini + fim) / 2

		if array[meio] == value {
			return meio
		} else if array[meio] < value {
			ini = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

func RecursiveBinarySearch(array []int, value int, ini int, fim int) int {
	if ini > fim {
		return -1
	}

	meio := (ini + fim) / 2

	if array[meio] == value {
		return meio
	} else if array[meio] < value {
		return RecursiveBinarySearch(array, value, meio+1, fim)
	} else {
		return RecursiveBinarySearch(array, value, ini, meio-1)
	}

}
