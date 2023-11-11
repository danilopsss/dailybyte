package main

import (
	"regexp"
)

type CounterUpper struct {
	upperCounted int
}

func (c *CounterUpper) addToCounter(multiplier int) {
	c.upperCounted += multiplier
}

func HasCorrectCapitalization(str string) bool {
	Counter := CounterUpper{0}
	allCaps := regexp.MustCompile(`[A-Z]`)

	for idx, char := range str {
		isCaps := allCaps.MatchString(string(char))
		if isCaps && idx > 0 && Counter.upperCounted != idx {
			return false
		} else if isCaps {
			Counter.addToCounter(1)
		}
	}
	return true
}
