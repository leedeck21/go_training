// Main entry point for Snakes and Ladders game
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"

	"my_golang_test_project/pkg/gameutils"
)

func main() {
	// We want to run the snake game multiple times concurrently to see which player wins most often.
	gamesPlayed := 1000000
	// Create a buffered channel to collect winners from each game
	winnersChannel := make(chan int, gamesPlayed)

	// Decide how many workers (goroutines) to run in parallel (usually # of CPU cores)
	numWorkers := runtime.NumCPU()

	// Create a jobs channel to send work (games to play) to the workers
	jobs := make(chan struct{}, gamesPlayed)

	// Start worker goroutines. Each worker will keep taking jobs from the jobs channel
	// and run playSnakeGame until the jobs channel is closed.
	for workerID := 0; workerID < numWorkers; workerID++ {
		go func() {
			for range jobs { // for each job received...
				playSnakeGame(winnersChannel) // ...run a game and send the result to winnersChannel
			}
		}()
	}

	// Send all jobs (one per game) into the jobs channel
	for i := 0; i < gamesPlayed; i++ {
		jobs <- struct{}{} // send an empty struct as a signal to play a game
	}
	close(jobs) // no more jobs to send; workers will finish when all jobs are done

	// Collect results from the winnersChannel as each game finishes
	winnersCount := make(map[int]int)
	for i := 0; i < gamesPlayed; i++ {
		winner := <-winnersChannel // receive the winner of a game
		winnersCount[winner]++     // count the win for that player
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

	// Print the percentage of games won by the top player
	if len(sortedWinners) > 0 {
		topPlayer := sortedWinners[0]
		percentage := float64(topPlayer.wins) / float64(gamesPlayed) * 100
		fmt.Printf("\nPlayer %d won the most games: %d times (%.2f%% of all games)\n", topPlayer.player, topPlayer.wins, percentage)
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