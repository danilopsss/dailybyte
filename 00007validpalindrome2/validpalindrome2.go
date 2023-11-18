package main


type String struct {
	word string
}

func (s String) hasEqualChar(start int, end int) bool {
	return s.word[start] == s.word[end] ||
		s.word[start + 1] == s.word[end] ||
		s.word[start] == s.word[end - 1]
   }

func (s String) isPalindrome() bool {
	for i := 1; i < len(s.word) - 2; i++ {
		if !s.hasEqualChar(i, len(s.word) - 1 - i) {
			return false
		}
	}
	return true
}