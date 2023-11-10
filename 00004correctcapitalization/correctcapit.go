package main

import (
	"regexp"
)

type Score struct {
	score int
}

func (s *Score) addToScore(multiplier int) {
	s.score += multiplier
}

func (s *Score) subtractToScore(multiplier int) {
	s.score -= multiplier
}

func HasCorrectCapitalization(str string) bool {
	Score := Score{0}
	allCaps := regexp.MustCompile(`[A-Z]`)

	for idx, char := range str {
		isCaps := allCaps.MatchString(string(char))
		if isCaps {
			Score.subtractToScore(1)
		} else if idx == 0 && isCaps {
			Score.addToScore(2)
		} else {
			Score.addToScore(1)
		}
	}
	sizeDiff := len(str) + int(Score.score)
	return sizeDiff % 2 == 0
}
