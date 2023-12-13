package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func tryAllPossibilitiesWithMemoization(spring string, damagedSpringGroupSizes []int) int {
	cacheString := spring + "," + strings.Join(utils.Map(damagedSpringGroupSizes, strconv.Itoa), ",")
	count, ok := cache[cacheString]
	if ok {
		return count
	}
	if len(damagedSpringGroupSizes) == 0 {
		for i := 0; i < len(spring); i++ {
			if spring[i] == '#' {
				cache[cacheString] = 0
				return 0
			}
		}
		cache[cacheString] = 1
		return 1
	}
	if len(spring) == 0 {
		cache[cacheString] = 0
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
			cache[cacheString] = tryAllPossibilitiesWithMemoization(spring[i+damagedSpringGroupSizes[0]+1:], damagedSpringGroupSizes[1:])
			return cache[cacheString]
		} else if spring[i] == '.' {
			cache[cacheString] = tryAllPossibilitiesWithMemoization(spring[i+1:], damagedSpringGroupSizes)
			return cache[cacheString]
		} else if spring[i] == '?' {
			cache[cacheString] = tryAllPossibilitiesWithMemoization("#"+spring[i+1:], damagedSpringGroupSizes) +
				tryAllPossibilitiesWithMemoization(spring[i+1:], damagedSpringGroupSizes)
			return cache[cacheString]
		}
	}
	return 0
}

func repeat(str string, times int, separator string) string {
	newStr := str
	for i := 1; i < times; i++ {
		newStr = newStr + separator + str
	}
	return newStr
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var sum = 0
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		spring := repeat(line[:strings.Index(line, " ")], 5, "?")
		damagedSpringGroupSizes := utils.Map(strings.Split(repeat(line[strings.Index(line, " ")+1:len(line)-1], 5, ","), ","), utils.StringToInt)
		sum += tryAllPossibilitiesWithMemoization(spring, damagedSpringGroupSizes)
	}
	fmt.Printf("%d\n", sum)
}
