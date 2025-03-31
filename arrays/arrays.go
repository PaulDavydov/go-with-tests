package arrays

func Sum(nums []int) int {
	sum := 0

	for _, i := range nums {
		sum += i
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, num := range nums {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int
	for _, i := range nums {
		if len(i) == 0 {
			sums = append(sums, 0)
		} else {
			tail := i[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
