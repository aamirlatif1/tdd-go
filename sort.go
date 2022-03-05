package tdd_go

func Sort(input []int) []int {
	if len(input) > 1 {
		for limit := 0; limit < len(input)-1; limit++ {
			for firstIndex := 0; firstIndex < len(input)-1; firstIndex++ {
				secondIndex := firstIndex + 1
				if input[firstIndex] > input[secondIndex] {
					temp := input[secondIndex]
					input[secondIndex] = input[firstIndex]
					input[firstIndex] = temp
				}
			}
		}
	}
	return input
}
