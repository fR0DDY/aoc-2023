package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"strings"
)

type SourceDestinationMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func readWhile(in *bufio.Reader, delimiter string) []SourceDestinationMap {
	var sourceDestinationMaps []SourceDestinationMap
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if line == delimiter {
			break
		}
		numbers := utils.Map(strings.Fields(line), utils.StringToInt)
		if len(numbers) > 0 {
			sourceDestinationMaps = append(sourceDestinationMaps, SourceDestinationMap{destinationRangeStart: numbers[0], sourceRangeStart: numbers[1], rangeLength: numbers[2]})
		}
	}
	return sourceDestinationMaps
}
