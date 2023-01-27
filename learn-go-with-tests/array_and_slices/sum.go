package array_and_slices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, sumTail(numbers))
	}

	return sums
}

func sumTail(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	tail := numbers[1:]
	return Sum(tail)
}
