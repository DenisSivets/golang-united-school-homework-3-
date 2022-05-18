package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var sl []int
	i := 0
	for k, _ := range input {
		sl[i] = k
		i++
	}

	sort.Ints(sl)
	i = 0
	for k, _ := range sl {
		result[i] = input[k]
		i++
	}

	return result
}
