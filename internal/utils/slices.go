package utils

func Unique(slice []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueSlice := []string{}

	for _, item := range slice {
		if _, found := uniqueMap[item]; !found {
			uniqueMap[item] = true
			uniqueSlice = append(uniqueSlice, item)
		}
	}

	return uniqueSlice
}
