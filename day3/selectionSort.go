package day3

func SelectionSort(input []int) []int {
	for i := len(input) - 1; i > 0; i-- {
		maxIndex := 0

		for j := 0; j <= i; j++ {
			if input[j] > input[maxIndex] {
				maxIndex = j
			}
		}

		input[i], input[maxIndex] = input[maxIndex], input[i]
	}
	return input
}
