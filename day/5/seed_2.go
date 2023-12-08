package main

import (
	"bufio"
	"com/adventofcode/2023/utils"
	"fmt"
	"os"
	"strings"
)

func getSourceValue(sourceDestinationMaps []SourceDestinationMap, destinationValue int) int {
	for _, sourceDestinationMap := range sourceDestinationMaps {
		if destinationValue >= sourceDestinationMap.destinationRangeStart && destinationValue < sourceDestinationMap.destinationRangeStart+sourceDestinationMap.rangeLength {
			return sourceDestinationMap.sourceRangeStart + (destinationValue - sourceDestinationMap.destinationRangeStart)
		}
	}
	return destinationValue
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

	foundLowestValue := false
	for locationValue := 0; !foundLowestValue; locationValue++ {
		humidityValue := getSourceValue(humidityToLocationMaps, locationValue)
		temperatureValue := getSourceValue(temperatureToHumidityMaps, humidityValue)
		lightValue := getSourceValue(lightToTemperatureMaps, temperatureValue)
		waterValue := getSourceValue(waterToLightMaps, lightValue)
		fertilizerValue := getSourceValue(fertilizerToWaterMaps, waterValue)
		soilValue := getSourceValue(soilToFertilizerMaps, fertilizerValue)
		seedValue := getSourceValue(seedToSoilMaps, soilValue)

		for i := 0; i < len(seeds); i += 2 {
			if seedValue > seeds[i] && seedValue < seeds[i]+seeds[i+1] {
				fmt.Printf("%d\n", locationValue)
				foundLowestValue = true
				break
			}
		}
	}
}
