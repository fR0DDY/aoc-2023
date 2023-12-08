package utils

import "strconv"

func StringToInt(item string) int {
	number, err := strconv.Atoi(item)
	if err == nil {
		return number
	}
	return 0
}

func StringToFloat(item string) float64 {
	number, err := strconv.ParseFloat(item, 8)
	if err == nil {
		return number
	}
	return 0
}
