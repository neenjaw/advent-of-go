package enum

import (
	"errors"
	"math"
)

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

func ContainsInt(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func Find(haystack []string, compare func(needle string) bool) (string, error) {
	for _, item := range haystack {
		if compare(item) {
			return item, nil
		}
	}
	return "", errors.New("nothing found with comparator")
}
