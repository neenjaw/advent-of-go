package enum

func EveryString(slice []string, fn func(string) bool) bool {
	result := true
	for _, s := range slice {
		result = result && fn(s)
	}
	return result
}
