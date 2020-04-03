package dynamic

func PowerSet(inSet []interface{}) [][]interface{} {
	powerSet := make([][]interface{}, 0)

	for _, elem := range inSet {
		for _, curSet := range powerSet {
			newSet := make([]interface{}, len(curSet))
			copy(newSet, curSet)
			newSet = append(newSet, elem)
			powerSet = append(powerSet, newSet)
		}
		curElemSet := make([]interface{}, 0, 1)
		curElemSet = append(curElemSet, elem)
		powerSet = append(powerSet, curElemSet)
	}
	powerSet = append(powerSet, make([]interface{}, 0))
	return powerSet
}
