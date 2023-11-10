package main

import (
	"strings"
	"regexp"
)

func ValidPalindrome(str string) bool {
	sanitizedStr := SanitizeString(str)
	for i := 0; i < len(sanitizedStr); i++ {
		if sanitizedStr[i] != sanitizedStr[len(sanitizedStr)-i-1] {
			return false
		}
	}
	return true
}

func SanitizeString(str string) string {
	patternRegex := regexp.MustCompile(`\W+`)

	cleanUpString:= func (p *regexp.Regexp, str string) string {
		s := p.ReplaceAll([]byte(str), []byte(""))
		return string(s)
	}

	return strings.ToLower(cleanUpString(patternRegex, str))
}
