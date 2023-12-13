package main

import (
	"com/adventofcode/2023/utils"
	"fmt"
)

func main() {
	var image []string
	for i := 0; ; i++ {
		var row string
		_, err := fmt.Scanf("%s", &row)
		if err != nil {
			break
		}
		image = append(image, row)
	}
	rows, cols := len(image), len(image[0])
	distance := utils.Matrix[int](len(image), len(image[0]))
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			distance[i][j] = 1
		}
	}
	for i := 0; i < rows; i++ {
		isEmpty := true
		for j := 0; j < cols; j++ {
			if image[i][j] != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			for j := 0; j < cols; j++ {
				distance[i][j] += 1
			}
		}
	}
	for i := 0; i < cols; i++ {
		isEmpty := true
		for j := 0; j < rows; j++ {
			if image[j][i] != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			for j := 0; j < rows; j++ {
				distance[j][i] += 1
			}
		}
	}
	var galaxyRows, galaxyCols []int
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == '#' {
				galaxyRows = append(galaxyRows, i)
				galaxyCols = append(galaxyCols, j)
			}
		}
	}
	sum := 0
	for i := 0; i < len(galaxyRows); i++ {
		for j := i + 1; j < len(galaxyRows); j++ {
			for m := min(galaxyRows[i], galaxyRows[j]); m < max(galaxyRows[i], galaxyRows[j]); m++ {
				sum += distance[m][galaxyCols[i]]
			}
			for m := min(galaxyCols[i], galaxyCols[j]); m < max(galaxyCols[i], galaxyCols[j]); m++ {
				sum += distance[galaxyRows[i]][m]
			}
		}
	}
	fmt.Printf("%d\n", sum)
}
