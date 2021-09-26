package enum

import "math"

func EveryString(slice []string, fn func(string) bool) bool {
	result := true
	for _, s := range slice {
		result = result && fn(s)
	}
	return result
}

func MaxByIntValue(values []int) (int, int) {
	maxIndex, maxValue := -1, math.MinInt
	for i, value := range values {
		if value > maxValue {
			maxIndex = i
			maxValue = value
		}
	}
	return maxIndex, maxValue
}

func Sum(values []int) (sum int) {
	for _, v := range values {
		sum += v
	}
	return
}
