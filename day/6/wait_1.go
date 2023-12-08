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
	times := utils.Map(strings.Fields(line[strings.Index(line, ":")+1:]), utils.StringToFloat)
	line, _ = in.ReadString('\n')
	distances := utils.Map(strings.Fields(line[strings.Index(line, ":")+1:]), utils.StringToFloat)

	product := 1
	for i := 0; i < len(times); i++ {
		d := times[i]*times[i] - 4.0*distances[i]
		root1 := (times[i] - math.Sqrt(d)) / (2.0)
		root2 := (times[i] + math.Sqrt(d)) / (2.0)
		multiple := int(math.Floor(root2-1e-9)) - int(math.Ceil(root1+1e-9)) + 1
		product *= multiple
	}
	fmt.Printf("%d\n", product)
}
