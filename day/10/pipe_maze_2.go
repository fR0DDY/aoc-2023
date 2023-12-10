package main

import (
	"com/adventofcode/2023/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	var maze []string
	startingRow, startingColumn := 0, 0
	for i := 0; ; i++ {
		var row string
		_, err := fmt.Scanf("%s", &row)
		if err != nil {
			break
		}
		maze = append(maze, row)
		if strings.Contains(row, "S") {
			startingRow, startingColumn = i, strings.Index(row, "S")
		}
	}
	rows, cols := len(maze), len(maze[0])
	type Move struct {
		rowMove   int
		colMove   int
		direction int
	}
	var pipeTypes = []uint8{'|', '-', 'L', 'J', '7', 'F'}
	const (
		NORTH = iota
		SOUTH
		EAST
		WEST
	)
	var initialDirection = [6]int{SOUTH, WEST, EAST, WEST, SOUTH, SOUTH}
	var south = [6]Move{{rowMove: -1, colMove: 0, direction: SOUTH}, {}, {}, {}, {rowMove: 0, colMove: -1, direction: EAST}, {rowMove: 0, colMove: +1, direction: WEST}}
	var north = [6]Move{{rowMove: 1, colMove: 0, direction: NORTH}, {}, {rowMove: 0, colMove: +1, direction: WEST}, {rowMove: 0, colMove: -1, direction: EAST}, {}, {}}
	var east = [6]Move{{}, {rowMove: 0, colMove: -1, direction: EAST}, {rowMove: -1, colMove: 0, direction: SOUTH}, {}, {}, {rowMove: +1, colMove: 0, direction: NORTH}}
	var west = [6]Move{{}, {rowMove: 0, colMove: +1, direction: WEST}, {}, {rowMove: -1, colMove: 0, direction: SOUTH}, {rowMove: +1, colMove: 0, direction: NORTH}, {}}
	for index, startingPipe := range pipeTypes {
		isVisited := utils.Matrix[uint8](rows, cols)
		// Using Shoelace formula to find area of loop (https://en.wikipedia.org/wiki/Shoelace_formula)
		currentRow, currentCol, currentPipe, direction, area := startingRow, startingColumn, startingPipe, initialDirection[index], 0
		isLoopCompleted := false
		loopLength := 0
		for ; ; loopLength++ {
			isVisited[currentRow][currentCol] = 1
			switch direction {
			case NORTH:
				pipeIndex := utils.SliceIndex[uint8](pipeTypes, currentPipe)
				if (Move{}) == north[pipeIndex] {
					currentRow = -1
				} else {
					area += (currentCol * (currentRow + north[pipeIndex].rowMove)) - ((currentCol + north[pipeIndex].colMove) * currentRow)
					currentRow = currentRow + north[pipeIndex].rowMove
					currentCol = currentCol + north[pipeIndex].colMove
					direction = north[pipeIndex].direction
				}

			case SOUTH:
				pipeIndex := utils.SliceIndex[uint8](pipeTypes, currentPipe)
				if (Move{}) == south[pipeIndex] {
					currentRow = -1
				} else {
					area += (currentCol * (currentRow + south[pipeIndex].rowMove)) - ((currentCol + south[pipeIndex].colMove) * currentRow)
					currentRow = currentRow + south[pipeIndex].rowMove
					currentCol = currentCol + south[pipeIndex].colMove
					direction = south[pipeIndex].direction
				}
			case EAST:
				pipeIndex := utils.SliceIndex[uint8](pipeTypes, currentPipe)
				if (Move{}) == east[pipeIndex] {
					currentRow = -1
				} else {
					area += (currentCol * (currentRow + east[pipeIndex].rowMove)) - ((currentCol + east[pipeIndex].colMove) * currentRow)
					currentRow = currentRow + east[pipeIndex].rowMove
					currentCol = currentCol + east[pipeIndex].colMove
					direction = east[pipeIndex].direction
				}
			case WEST:
				pipeIndex := utils.SliceIndex[uint8](pipeTypes, currentPipe)
				if (Move{}) == west[pipeIndex] {
					currentRow = -1
				} else {
					area += (currentCol * (currentRow + west[pipeIndex].rowMove)) - ((currentCol + west[pipeIndex].colMove) * currentRow)
					currentRow = currentRow + west[pipeIndex].rowMove
					currentCol = currentCol + west[pipeIndex].colMove
					direction = west[pipeIndex].direction
				}
			}
			if currentRow < 0 || currentRow >= rows || currentCol < 0 || currentCol >= cols {
				break
			}
			if maze[currentRow][currentCol] == '.' {
				break
			}
			if currentRow == startingRow && currentCol == startingColumn {
				if initialDirection[index] == direction {
					isLoopCompleted = true
				}
				break
			}
			currentPipe = maze[currentRow][currentCol]
		}
		if isLoopCompleted {
			// Answer is area - loop length - 1
			fmt.Printf("%d\n", int(math.Abs(float64(area)/2.0)-math.Floor(float64(loopLength)/2.0)))
		}
	}
}
