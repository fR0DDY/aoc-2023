package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		colonIndex := strings.Index(line, ":")
		number, _ := strconv.Atoi(line[5:colonIndex])
		sets := line[colonIndex+1 : len(line)-1]
		isPossible := true
		for _, set := range strings.Split(sets, ";") {
			for _, cubes := range strings.Split(set, ",") {
				var value int
				var color string
				fmt.Sscanf(cubes, "%d %s", &value, &color)
				if color == "red" && value > 12 {
					isPossible = false
					break
				} else if color == "green" && value > 13 {
					isPossible = false
					break
				} else if color == "blue" && value > 14 {
					isPossible = false
					break
				}
			}
		}
		if isPossible {
			sum += number
		}
	}
	fmt.Printf("%d\n", sum)
}
