package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var instruction string
	in := bufio.NewReader(os.Stdin)
	fmt.Scanf("%s\n\n", &instruction)

	type Node struct {
		left  string
		right string
	}
	network := make(map[string]Node)
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		source := line[0:strings.Index(line, " ")]
		network[source] = Node{left: line[strings.Index(line, "(")+1 : strings.Index(line, ",")], right: line[strings.Index(line, ",")+2 : strings.LastIndex(line, ")")]}
	}
	instructionLength := len(instruction)
	steps := 0
	for i, currentNode := 0, "AAA"; ; i, steps = (i+1)%instructionLength, steps+1 {
		if currentNode == "ZZZ" {
			break
		}
		if instruction[i] == 'L' {
			currentNode = network[currentNode].left
		} else {
			currentNode = network[currentNode].right
		}
	}
	fmt.Printf("%d\n", steps)

}
