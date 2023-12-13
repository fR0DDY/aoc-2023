package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"os"
	"strings"
)

func tryAllPossibilities(spring string, damagedSpringGroupSizes []int) int {
	if len(damagedSpringGroupSizes) == 0 {
		for i := 0; i < len(spring); i++ {
			if spring[i] == '#' {
				return 0
			}
		}
		return 1
	}
	if len(spring) == 0 {
		return 0
	}
	for i := 0; i < len(spring); i++ {
		if spring[i] == '#' {
			for k := i; k < i+damagedSpringGroupSizes[0]; k++ {
				if k >= len(spring) {
					return 0
				}
				if spring[k] == '.' {
					return 0
				}
			}
			if i+damagedSpringGroupSizes[0] < len(spring) && spring[i+damagedSpringGroupSizes[0]] == '#' {
				return 0
			} else if i+damagedSpringGroupSizes[0] == len(spring) {
				if len(damagedSpringGroupSizes) == 1 {
					return 1
				} else {
					return 0
				}
			}
			return tryAllPossibilities(spring[i+damagedSpringGroupSizes[0]+1:], damagedSpringGroupSizes[1:])
		} else if spring[i] == '.' {
			return tryAllPossibilities(spring[i+1:], damagedSpringGroupSizes)
		} else if spring[i] == '?' {
			return tryAllPossibilities("#"+spring[i+1:], damagedSpringGroupSizes) +
				tryAllPossibilities(spring[i+1:], damagedSpringGroupSizes)
		}
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var sum = 0
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		spring := line[:strings.Index(line, " ")]
		damagedSpringGroupSizes := utils.Map(strings.Split(line[strings.Index(line, " ")+1:len(line)-1], ","), utils.StringToInt)
		sum += tryAllPossibilities(spring, damagedSpringGroupSizes)
	}
	fmt.Printf("%d\n", sum)
}
