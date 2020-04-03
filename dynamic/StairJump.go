package dynamic

func StairJumpCount(stairsCount int, jumps []int) int {
	cache := make([]int, stairsCount+1, stairsCount+1)
	return stairJumpCount(stairsCount, jumps, cache)
}

func stairJumpCount(stairsCount int, jumps []int, cache []int) int {
	if stairsCount < 0 {
		return 0
	}
	if stairsCount == 0 {
		return 1
	}

	var totalWays int
	for _, curJump := range jumps {
		curJumpIdx := stairsCount - curJump
		if curJumpIdx < 0 {
			continue
		}
		curWays := cache[curJumpIdx]

		if curWays == 0 {
			curWays = stairJumpCount(curJumpIdx, jumps, cache)
			cache[curJumpIdx] = curWays
		}
		totalWays += curWays
	}
	return totalWays
}
