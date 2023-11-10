package main

import (
	"math"
	"regexp"
)

type Score struct {
	score float64
}

func (s *Score) addToScore(multiplier int) {
	s.score += float64(multiplier)
}

func (s *Score) subtractToScore(multiplier int) {
	s.score -= float64(multiplier)
}

func (s *Score) getScore() int {
	absoluteScore := math.Abs(s.score)
	return int(absoluteScore)
}

func HasCorrectCapitalization(str string) bool {
	Score := Score{0}
	allCaps := regexp.MustCompile(`[a-z]`)

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
	sizeDiff := len(str) - Score.getScore()
	return sizeDiff % 2 == 0
}
