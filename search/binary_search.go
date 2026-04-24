package search

func IterativeBinarySearch(v []int, value int) int {
	ini := 0
	fim := len(v) - 1

	for ini <= fim {
		meio := (ini + fim) / 2

		if v[meio] == value {
			return meio
		} else if v[meio] < value {
			ini = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

func RecursiveBinarySearch(v []int, value int, ini int, fim int) int {
	if ini > fim {
		return -1
	}

	meio := (ini + fim) / 2

	if v[meio] == value {
		return meio
	} else if v[meio] < value {
		return RecursiveBinarySearch(v, value, meio+1, fim)
	} else {
		return RecursiveBinarySearch(v, value, ini, meio-1)
	}
}
