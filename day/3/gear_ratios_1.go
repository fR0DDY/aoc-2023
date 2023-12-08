package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {

	var schematic []string
	var sum = 0
	for {
		var row string
		n, err := fmt.Scanf("%s\n", &row)
		if n == 0 || err != nil {
			break
		}
		schematic = append(schematic, row)
	}
	rows := len(schematic)
	columns := len(schematic[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if unicode.IsDigit(rune(schematic[i][j])) {
				start := j
				for j < columns && unicode.IsDigit(rune(schematic[i][j])) {
					j++
				}
				end := j
				isPartNumber := false
				if i-1 >= 0 {
					for k := max(0, start-1); k < min(columns-1, end+1); k++ {
						if schematic[i-1][k] != '.' {
							isPartNumber = true
						}
					}
				}
				if !isPartNumber && i+1 < rows {
					for k := max(0, start-1); k < min(columns-1, end+1); k++ {
						if schematic[i+1][k] != '.' {
							isPartNumber = true
						}
					}
				}
				if !isPartNumber && start-1 >= 0 && schematic[i][start-1] != '.' {
					isPartNumber = true
				}
				if !isPartNumber && end < columns && schematic[i][end] != '.' {
					isPartNumber = true
				}
				if isPartNumber {
					number, err := strconv.Atoi(schematic[i][start:end])
					if err == nil {
						sum += number
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", sum)
}
