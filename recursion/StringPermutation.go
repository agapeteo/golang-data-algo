package recursion

func PermuteStr(str string) []string {
	perms := make([]string, 0)

	return permuteStr([]rune(""), []rune(str), &perms)
}

func permuteStr(prefix, str []rune, perms *[]string) []string {
	if len(str) > 0 {
		for idx, char := range str {
			newPrefix := append(prefix, char)
			strBeforeRune := str[:idx]
			strAfterRune := str[idx+1:]
			newStr := append(strAfterRune, strBeforeRune...)

			permuteStr(newPrefix, newStr, perms)
		}
	} else {
		*perms = append(*perms, string(prefix))
	}
	return *perms
}
