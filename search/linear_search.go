package search

func LinearSearch(v []int, value int) int {
	for i := 0; i < len(v); i++ {
		if v[i] == value {
			return i
		}
	}
	return -1
}
