package bits

func PowerSet(inSet []interface{}) [][]interface{} {
	powerSet := make([][]interface{}, 0)

	max := 1 << uint(len(inSet))
	for i := 0; i < max; i++ {
		powerSet = append(powerSet, setForNUmber(i, inSet))
	}
	return powerSet
}

func setForNUmber(n int, inSet []interface{}) []interface{} {
	set := make([]interface{}, 0)

	i := 0
	num := n

	for num > 0 {
		if (num & 1) != 0 {
			set = append(set, inSet[i])
		}
		num = num >> 1
		i++
	}
	return set
}
