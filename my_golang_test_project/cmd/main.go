// Main entry point for Snakes and Ladders game
package main

import (
	"fmt"
	"math/rand"

	"my_golang_test_project/pkg/gameutils"
)

func main() {
	// We want to run the snake game multiple times concurrently to see which player wins most often.

	playSnakeGame("Lee")

	// I need a snake game
	// A way to run multiple games at the same time
	// A place to store results
	// A place to display results
}

func diceRoll() int {
	return rand.Intn(6) + 1
}

// isSnake and isLadder moved to gameutils package

func playSnakeGame(playerName string) {
	playerPosition := 0

	for i := 0; i < 100; i++ {
		playerPosition = takeTurn(playerName, playerPosition)
	}

	fmt.Printf("Player position is %d", playerPosition)

}

func takeTurn(playerName string, currentPosition int) int {
	diceRollOutcome := diceRoll()
	currentPosition += diceRollOutcome
	currentPosition = gameutils.IsSnake(currentPosition)
	currentPosition = gameutils.IsLadder(currentPosition)

	return currentPosition
}