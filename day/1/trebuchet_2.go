package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var amendedValue string
	var sum = 0
	for {
		n, err := fmt.Scanf("%s\n", &amendedValue)
		if n == 0 || err != nil {
			break
		}
		amendedValue = strings.ReplaceAll(amendedValue, "one", "o1e")
		amendedValue = strings.ReplaceAll(amendedValue, "two", "t2o")
		amendedValue = strings.ReplaceAll(amendedValue, "three", "t3e")
		amendedValue = strings.ReplaceAll(amendedValue, "four", "4")
		amendedValue = strings.ReplaceAll(amendedValue, "five", "5e")
		amendedValue = strings.ReplaceAll(amendedValue, "six", "6")
		amendedValue = strings.ReplaceAll(amendedValue, "seven", "7")
		amendedValue = strings.ReplaceAll(amendedValue, "eight", "e8t")
		amendedValue = strings.ReplaceAll(amendedValue, "nine", "9e")
		digits := regexp.MustCompile("\\D+").ReplaceAllLiteralString(amendedValue, "")
		actualValue := string(digits[0]) + string(digits[len(digits)-1])
		number, err := strconv.Atoi(actualValue)
		if err == nil {
			sum += number
		}
	}
	fmt.Printf("%d\n", sum)
}
