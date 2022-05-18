package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var sl []int

	for k, _ := range input {
		sl = append(sl, k)
	}

	sort.Ints(sl)

	for _, k := range sl {
		result = append(result, input[k])
	}

	return result
}
