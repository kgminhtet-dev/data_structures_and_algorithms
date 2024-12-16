package day3

func BubbleSort(input []int) []int {
	if len(input) == 0 || len(input) == 1 {
		return input
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input)-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}

	return input
}
