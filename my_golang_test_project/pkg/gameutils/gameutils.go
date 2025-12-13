package gameutils

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
