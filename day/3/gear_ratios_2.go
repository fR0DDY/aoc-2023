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
	type NumbersWithLocation struct {
		number int
		start  int
		end    int
	}
	var numberArray [][]NumbersWithLocation
	for i := 0; i < rows; i++ {
		var rowNumbers []NumbersWithLocation
		for j := 0; j < columns; j++ {
			if unicode.IsDigit(rune(schematic[i][j])) {
				start := j
				for j < columns && unicode.IsDigit(rune(schematic[i][j])) {
					j++
				}
				end := j
				number, err := strconv.Atoi(schematic[i][start:end])
				if err == nil {
					numberStruct := NumbersWithLocation{number: number, start: start, end: end - 1}
					rowNumbers = append(rowNumbers, numberStruct)
				}
			}
		}
		numberArray = append(numberArray, rowNumbers)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if schematic[i][j] == '*' {
				count := 0
				gearRatio := 1
				for k := max(0, i-1); k < min(i+2, rows); k++ {
					for _, numberStruct := range numberArray[k] {
						if (numberStruct.start > j-2 && numberStruct.start < j+2) || (numberStruct.end > j-2 && numberStruct.end < j+2) {
							count++
							gearRatio *= numberStruct.number
						}
					}
				}
				if count == 2 {
					sum += gearRatio
				}
			}
		}
	}
	fmt.Printf("%d\n", sum)
}
