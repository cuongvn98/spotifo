package utils

func IndexOfInStringSlice(sl []string, el string) int {
	for i, v := range sl {
		if v == el {
			return i
		}
	}

	return -1
}
