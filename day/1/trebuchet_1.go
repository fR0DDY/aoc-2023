package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	var amendedValue string
	var sum = 0
	for {
		n, err := fmt.Scanf("%s\n", &amendedValue)
		if n == 0 || err != nil {
			break
		}
		digits := regexp.MustCompile("\\D+").ReplaceAllLiteralString(amendedValue, "")
		actualValue := string(digits[0]) + string(digits[len(digits)-1])
		number, err := strconv.Atoi(actualValue)
		if err == nil {
			sum += number
		}
	}
	fmt.Printf("%d\n", sum)
}
