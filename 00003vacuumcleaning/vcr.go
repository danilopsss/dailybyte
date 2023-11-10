package main

const (
	right, up = +1, +1
	left, down = -1, -1
)

func MarkMove(char rune) int {
	switch char {
	case 'R':
		return right
	case 'L':
		return left
	case 'U':
		return up
	case 'D':
		return down
	default:
		return 0
	}
}

func StoppedAtAStartingPos(str string) bool {
	var position int
	for _, char := range str {
		position += MarkMove(char)
	}
	return position == 0
}
