package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var winningCards [][]int
	var myCards [][]int
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		winningNumbers := utils.Map(strings.Fields(line[strings.Index(line, ":")+1:strings.Index(line, "|")]), utils.StringToInt)
		myNumbers := utils.Map(strings.Fields(line[strings.Index(line, "|")+1:]), utils.StringToInt)
		winningCards = append(winningCards, winningNumbers)
		myCards = append(myCards, myNumbers)
	}
	cards := len(winningCards)
	instances := make([]int, cards)
	instances = utils.Map(instances, func(item int) int { return 1 })
	for i := 0; i < cards; i++ {
		count := 0
		for _, number := range myCards[i] {
			if slices.Contains(winningCards[i], number) {
				count += 1
			}
		}
		for j := i + 1; j < i+count+1; j++ {
			instances[j] += instances[i]
		}
	}
	sum := utils.Reduce(instances, func(accum int, value int) int { return accum + value })
	fmt.Printf("%d\n", sum)
}
