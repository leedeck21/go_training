// Main entry point for Snakes and Ladders game
package main

import (
	"math/rand"
	"my_golang_test_project/pkg/gameutils"
	"runtime"
	"time"
)

func main() {
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
			go func(wid int) {
				r := rand.New(rand.NewSource(int64(wid) + time.Now().UnixNano()))
				playerPositions := map[int]int{
					1: 0,
					2: 0,
					3: 0,
					4: 0,
				}
				for range jobs {
					for k := range playerPositions {
						playerPositions[k] = 0
					}
					playSnakeGame(winnersChannel, r, playerPositions)
				}
			}(workerID)
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

	   sortedWinners := gameutils.AggregateAndSortWinners(winnersCount)
	   gameutils.PrintPlayerWinStatistics(sortedWinners, gamesPlayed)
}


// diceRoll now uses the provided rand.Rand
func diceRoll(r *rand.Rand) int {
	return r.Intn(6) + 1
}

// isSnake and isLadder moved to gameutils package


func playSnakeGame(winnersChannel chan int, r *rand.Rand, playerPositions map[int]int) {
       winner := 0

       for winner == 0 {
	       for playerNumber, playerPosition := range playerPositions {
		       newPosition := takeTurn(playerPosition, r)
		       playerPositions[playerNumber] = newPosition

		       if newPosition == 100 {
			       winner = playerNumber
			       break
		       }
	       }
       }
       winnersChannel <- winner
}


func takeTurn(currentPosition int, r *rand.Rand) int {
	diceRollOutcome := diceRoll(r)
	newPosition := diceRollOutcome + currentPosition

	if newPosition > 100 {
		return currentPosition
	}

	newPosition = gameutils.IsSnake(newPosition)
	newPosition = gameutils.IsLadder(newPosition)

	return newPosition
}