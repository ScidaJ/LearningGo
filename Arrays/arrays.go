package arrays

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numsToSums ...[]int) []int {
	var sums []int

	for _, nums := range numsToSums {
		sums = append(sums, Sum(nums))
	}

	return sums

}

func SumAllTails(numsToSums ...[]int) []int {
	var sums []int

	for _, nums := range numsToSums {
		numsLen := len(nums)
		tail := nums[(numsLen - 1):]
		sums = append(sums, Sum(tail))
	}

	return sums
}
