package main

import (
	"bufio"
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
		colonIndex := strings.Index(line, ":")
		sets := line[colonIndex+1 : len(line)-1]
		minRedBall, minBlueBall, minGreenBall := 0, 0, 0
		for _, set := range strings.Split(sets, ";") {
			for _, cubes := range strings.Split(set, ",") {
				var value int
				var color string
				fmt.Sscanf(cubes, "%d %s", &value, &color)
				if color == "red" && value > minRedBall {
					minRedBall = value
				} else if color == "green" && value > minGreenBall {
					minGreenBall = value
				} else if color == "blue" && value > minBlueBall {
					minBlueBall = value
				}
			}
		}
		sum += minRedBall * minBlueBall * minGreenBall
	}
	fmt.Printf("%d\n", sum)
}
