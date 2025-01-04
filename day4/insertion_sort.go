package day4

func InsertionSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		current := input[i]
		j := i - 1
		for ; j >= 0 && current < input[j]; j-- {
			input[j+1] = input[j]
		}
		input[j+1] = current
	}
	return input
}
