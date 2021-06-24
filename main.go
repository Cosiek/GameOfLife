package main

import (
	"time"
)

type Coordinates struct {
	x int // I assume int is a good enough aproximation for infinity
	y int
}

func GameOfLife(liveCells map[Coordinates]bool) {
	// play
	for i := 0; true; i++ {
		printState(liveCells, i)
		if len(liveCells) == 0 {
			break
		}
		liveCells = Tick(liveCells)
		time.Sleep(500 * time.Millisecond)
	}
}

func Tick(nowLiveCells map[Coordinates]bool) map[Coordinates]bool {
	// gather a set of cells that require consideration
	toBeConsidered := make(map[Coordinates]bool)
	for baseCellCoords, _ := range nowLiveCells {
		// iterate base cell neighbours
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				consideredCoords := Coordinates{
					baseCellCoords.x + x,
					baseCellCoords.y + y,
				}
				toBeConsidered[consideredCoords] = true // doesn't matter if it's true
			}
		}
	}
	// iterate cells that need to be considered
	liveCells := make(map[Coordinates]bool)
	for coordinates, _ := range toBeConsidered {
		liveNeighboursCounter := 0
		_, isAlive := nowLiveCells[coordinates]
		// iterate their neighbours
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				// a cell isn't its own neighbour
				if x == 0 && y == 0 {
					continue
				}
				consideredCoords := Coordinates{
					coordinates.x + x,
					coordinates.y + y,
				}

				if _, ok := nowLiveCells[consideredCoords]; ok {
					liveNeighboursCounter++
				}

			}
		}
		// check if alive
		if (isAlive && liveNeighboursCounter == 2) || liveNeighboursCounter == 3 {
			liveCells[coordinates] = true
		}
	}

	return liveCells
}

// PRINTING

func printState(liveCells map[Coordinates]bool, tickNo int) {
	// clear screen
	println("\033[H\033[2J")
	// start printing
	gridSize := 25
	println("Tick:", tickNo)
	for y := 0; y <= gridSize; y++ {
		print("|")
		for x := 0; x <= gridSize; x++ {
			coordinates := Coordinates{x, y}
			if _, ok := liveCells[coordinates]; ok {
				print("X")
			} else {
				print(" ")
			}
		}
		println("|")
	}
}

// MAIN

func main() {
	// init data structure
	liveCells := make(map[Coordinates]bool)
	// set initial state - the glider
	liveCells[Coordinates{12, 12}] = true
	liveCells[Coordinates{13, 12}] = true
	liveCells[Coordinates{13, 14}] = true
	liveCells[Coordinates{14, 12}] = true
	liveCells[Coordinates{14, 13}] = true
	// add another shape on gliders way
	liveCells[Coordinates{20, 2}] = true
	liveCells[Coordinates{20, 3}] = true
	liveCells[Coordinates{20, 4}] = true
	// run the game
	GameOfLife(liveCells)
}
