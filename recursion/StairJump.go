package recursion

func StairJumpCount(stairsCount int, jumps []int) int {
	if stairsCount < 0 {
		return 0
	}

	if stairsCount == 0 {
		return 1
	}

	var total int
	for _, curJumps := range jumps {
		total += StairJumpCount(stairsCount-curJumps, jumps)
	}
	return total
}
