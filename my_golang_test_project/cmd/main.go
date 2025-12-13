// Main entry point for Snakes and Ladders game
package main

import (
	"fmt"
	"math/rand"
	"sort"

	"my_golang_test_project/pkg/gameutils"
)

func main() {
	// We want to run the snake game multiple times concurrently to see which player wins most often.
	winnersChannel := make(chan int)

	gamesPlayed := 1000000

	for i := 0; i < gamesPlayed; i++ {
		go playSnakeGame(winnersChannel)
	}

	winnersCount := make(map[int]int)

	for i := 0; i < gamesPlayed; i++ {
		winner := <-winnersChannel
		winnersCount[winner]++
	}

	// Convert map to slice of structs for sorting
	type playerWin struct {
		player int
		wins   int
	}

	var sortedWinners []playerWin

	for player, wins := range winnersCount {
		sortedWinners = append(sortedWinners, playerWin{player, wins})
	}

	// Sort by wins descending
	sort.Slice(sortedWinners, func(left, right int) bool {
		return sortedWinners[left].wins > sortedWinners[right].wins
	})

	for _, pw := range sortedWinners {
		fmt.Printf("Player %d won %d times\n", pw.player, pw.wins)
	}
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