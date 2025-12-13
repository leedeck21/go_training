package gameutils

import (
	"fmt"
	"sort"
)

// AggregateAndSortWinners converts a map of player win counts to a sorted slice of PlayerWin (descending by wins)
func AggregateAndSortWinners(winnersCount map[int]int) []PlayerWin {
       var sortedWinners []PlayerWin
       for player, wins := range winnersCount {
	       sortedWinners = append(sortedWinners, PlayerWin{Player: player, Wins: wins})
       }
       // Sort by wins descending
       sort.Slice(sortedWinners, func(left, right int) bool {
	       return sortedWinners[left].Wins > sortedWinners[right].Wins
       })
       return sortedWinners
}

// PlayerWin holds the win count for a player
type PlayerWin struct {
	Player int
	Wins   int
}

// PrintPlayerWinStatistics prints each player's win count, percentage, and compares the top player to others
func PrintPlayerWinStatistics(sortedWinners []PlayerWin, gamesPlayed int) {
       fmt.Println("\nPlayer win statistics:")
       for _, pw := range sortedWinners {
	       percent := float64(pw.Wins) / float64(gamesPlayed) * 100
	       fmt.Printf("Player %d won %d times (%.2f%%)\n", pw.Player, pw.Wins, percent)
       }

       if len(sortedWinners) > 0 {
	       topPlayer := sortedWinners[0]
	       topPercent := float64(topPlayer.Wins) / float64(gamesPlayed) * 100
	       fmt.Printf("\nPlayer %d won the most games: %d times (%.2f%% of all games)\n", topPlayer.Player, topPlayer.Wins, topPercent)

	       for i := 1; i < len(sortedWinners); i++ {
		       other := sortedWinners[i]
		       otherPercent := float64(other.Wins) / float64(gamesPlayed) * 100
		       if otherPercent > 0 {
			       percentDiff := topPercent - otherPercent
			       fmt.Printf("Player %d wins %.2f%% more often than Player %d\n", topPlayer.Player, percentDiff, other.Player)
		       } else {
			       fmt.Printf("Player %d never won, so Player %d is infinitely more likely to win.\n", other.Player, topPlayer.Player)
		       }
	       }
       }
}

func IsSnake(position int) int {
	snakePositions := map[int]int{
		32: 10,
		36: 6,
		48: 26,
		62: 18,
		88: 24,
		95: 56,
		97: 78,
	}

	snakePosition, exists := snakePositions[position]

	if exists {
		return snakePosition
	}

	return position
}

func IsLadder(position int) int {
	ladderPositions := map[int]int{
		1:  38,
		4:  14,
		8:  10,
		21: 42,
		28: 78,
		50: 67,
		88: 99,
		71: 92,
	}

	ladderPosition, exists := ladderPositions[position]

	if exists {
		return ladderPosition
	}

	return position
}
