package array_and_slices

func Sum(nums [5]int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
