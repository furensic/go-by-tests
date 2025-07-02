package array_and_slices

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sum := []int{}
	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}
	return sum
}
