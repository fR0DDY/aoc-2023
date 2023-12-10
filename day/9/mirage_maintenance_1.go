package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var sum = 0
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		var numbers [][]int
		numbers = append(numbers, utils.Map(strings.Fields(line), utils.StringToInt))
		for i, isAllZero := 0, true; ; i, isAllZero = i+1, true {
			numbers = append(numbers, make([]int, 0))
			for j := 1; j < len(numbers[i]); j++ {
				diff := numbers[i][j] - numbers[i][j-1]
				numbers[i+1] = append(numbers[i+1], diff)
				if diff != 0 {
					isAllZero = false
				}
			}
			if isAllZero {
				break
			}
		}
		nextValue := 0
		for i := len(numbers) - 1; i >= 0; i-- {
			nextValue += numbers[i][len(numbers[i])-1]
		}
		sum += nextValue
	}

	fmt.Printf("%d\n", sum)
}
