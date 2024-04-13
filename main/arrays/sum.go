package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var ans []int
	for _, arr := range numbersToSum {
		ans = append(ans, Sum(arr))
	}
	return ans
}
