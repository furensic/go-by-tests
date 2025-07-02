package array_and_slices

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numbersOfSlices := len(numbersToSum)
	sum := make([]int, numbersOfSlices)

	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	numbersOfSlices := len(numbersToSum)
	sum := make([]int, numbersOfSlices)
	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers[1:])
	}

	return sum
}
