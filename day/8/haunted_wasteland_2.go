package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
	if len(integers) < 2 {
		return integers[0]
	}
	result := 1

	for i := 0; i < len(integers); i++ {
		result = result * integers[i] / GCD(result, integers[i])
	}

	return result
}

func main() {
	var instruction string
	in := bufio.NewReader(os.Stdin)
	fmt.Scanf("%s\n\n", &instruction)

	type Node struct {
		left  string
		right string
	}
	network := make(map[string]Node)
	var sourceNodes []string

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		source := line[0:strings.Index(line, " ")]
		network[source] = Node{left: line[strings.Index(line, "(")+1 : strings.Index(line, ",")], right: line[strings.Index(line, ",")+2 : strings.LastIndex(line, ")")]}
		if source[2] == 'A' {
			sourceNodes = append(sourceNodes, source)
		}
	}
	instructionLength := len(instruction)
	var periods []int
	for j := 0; j < len(sourceNodes); j++ {
		steps := 0
		for i := 0; ; i, steps = (i+1)%instructionLength, steps+1 {
			if sourceNodes[j][2] == 'Z' {
				periods = append(periods, steps)
				break
			}
			if instruction[i] == 'L' {
				sourceNodes[j] = network[sourceNodes[j]].left
			} else {
				sourceNodes[j] = network[sourceNodes[j]].right
			}

		}
	}
	fmt.Printf("%d\n", LCM(periods))

}
