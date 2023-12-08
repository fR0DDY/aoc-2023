package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		winningNumbers := utils.Map(strings.Fields(line[strings.Index(line, ":")+1:strings.Index(line, "|")]), utils.StringToInt)
		myNumbers := utils.Map(strings.Fields(line[strings.Index(line, "|")+1:]), utils.StringToInt)
		count := 0
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				count += 1
			}
		}
		if count > 0 {
			sum += int(math.Pow(2.0, float64(count-1)))
		}
	}
	fmt.Printf("%d\n", sum)
}
