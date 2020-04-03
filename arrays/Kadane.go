package arrays

func maxSum(s []int) int {
	result := 0
	curSum := 0

	for i, _ := range s {
		if curSum+s[i] > 0 {
			curSum += s[i]
			if curSum > result {
				result = curSum
			}
		} else {
			curSum = 0
		}
	}
	return result
}
