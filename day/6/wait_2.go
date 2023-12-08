package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	time := utils.StringToFloat(strings.ReplaceAll(line[strings.Index(line, ":")+1:len(line)-1], " ", ""))
	line, _ = in.ReadString('\n')
	distance := utils.StringToFloat(strings.ReplaceAll(line[strings.Index(line, ":")+1:], " ", ""))
	d := time*time - 4.0*distance
	root1 := (time - math.Sqrt(d)) / (2.0)
	root2 := (time + math.Sqrt(d)) / (2.0)
	multiple := int(math.Floor(root2-1e-9)) - int(math.Ceil(root1+1e-9)) + 1
	fmt.Printf("%d\n", multiple)
}
