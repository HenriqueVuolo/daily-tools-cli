package utils

func RemoveEmptyStrings(list []string) []string {
	var result []string

	for _, str := range list {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}
