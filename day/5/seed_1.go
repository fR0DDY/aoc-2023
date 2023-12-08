package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func getDestinationValue(sourceDestinationMaps []SourceDestinationMap, sourceValue int) int {
	for _, sourceDestinationMap := range sourceDestinationMaps {
		if sourceValue >= sourceDestinationMap.sourceRangeStart && sourceValue < sourceDestinationMap.sourceRangeStart+sourceDestinationMap.rangeLength {
			return sourceDestinationMap.destinationRangeStart + (sourceValue - sourceDestinationMap.sourceRangeStart)
		}
	}
	return sourceValue
}

func main() {
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	if err != nil {
		panic("Unexpected Input")
	}
	seeds := utils.Map(strings.Fields(line[strings.Index(line, ":")+1:]), utils.StringToInt)
	readWhile(in, "seed-to-soil map:\n")
	seedToSoilMaps := readWhile(in, "soil-to-fertilizer map:\n")
	soilToFertilizerMaps := readWhile(in, "fertilizer-to-water map:\n")
	fertilizerToWaterMaps := readWhile(in, "water-to-light map:\n")
	waterToLightMaps := readWhile(in, "light-to-temperature map:\n")
	lightToTemperatureMaps := readWhile(in, "temperature-to-humidity map:\n")
	temperatureToHumidityMaps := readWhile(in, "humidity-to-location map:\n")
	humidityToLocationMaps := readWhile(in, "humidity-to-location map:\n")

	lowestLocation := math.MaxInt
	for _, seedValue := range seeds {
		soilValue := getDestinationValue(seedToSoilMaps, seedValue)
		fertilizerValue := getDestinationValue(soilToFertilizerMaps, soilValue)
		waterValue := getDestinationValue(fertilizerToWaterMaps, fertilizerValue)
		lightValue := getDestinationValue(waterToLightMaps, waterValue)
		temperatureValue := getDestinationValue(lightToTemperatureMaps, lightValue)
		humidityValue := getDestinationValue(temperatureToHumidityMaps, temperatureValue)
		locationValue := getDestinationValue(humidityToLocationMaps, humidityValue)
		if locationValue < lowestLocation {
			lowestLocation = locationValue
		}
	}
	fmt.Printf("%d\n", lowestLocation)
}
