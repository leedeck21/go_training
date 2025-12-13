// Main entry point for Snakes and Ladders game
package main

import (
	"fmt"
	"math/rand"

	"my_golang_test_project/pkg/gameutils"
)

func main() {
	// We want to run the snake game multiple times concurrently to see which player wins most often.
	winnersChannel := make(chan int)

	for i := 0; i < 100; i++ {
		fmt.Printf("Starting game %d\n", i+1)
		go playSnakeGame(winnersChannel)
		fmt.Println()
	}

	// I need a snake game
	// A way to run multiple games at the same time
	// A place to store results
	// A place to display results
}

func diceRoll() int {
	return rand.Intn(6) + 1
}

// isSnake and isLadder moved to gameutils package

func playSnakeGame(winnersChannel chan int) {
	playerPositions := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}

	winner := 0

	for winner == 0 {
		for playerNumber, playerPosition := range playerPositions {
			newPosition := takeTurn(playerPosition)
			playerPositions[playerNumber] = newPosition

			if newPosition == 100 {
				winner = playerNumber

				break
			}
		}
	}

	fmt.Printf("Game winner is %d", winner)
	winnersChannel <- winner
}

func takeTurn(currentPosition int) int {
	diceRollOutcome := diceRoll()
	newPosition := diceRollOutcome + currentPosition

	if newPosition > 100 {
		return currentPosition
	}

	newPosition = gameutils.IsSnake(newPosition)
	newPosition = gameutils.IsLadder(newPosition)

	return newPosition
}