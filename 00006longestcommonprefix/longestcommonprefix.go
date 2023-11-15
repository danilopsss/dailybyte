package main

type stringEval struct {
	wordList []string
}

func (s stringEval) LoopSize() (int, int) {
	min := int(^uint(0) >> 1)
	var index int
	for idx, word := range s.wordList {
		if len(word) < min {
			min = len(word)
			index = idx
		}
	}
	return index, min
}

func (s stringEval) IndexIsValid(i int) bool {
	return s.wordList[0][i] == s.wordList[1][i] && s.wordList[1][i] == s.wordList[2][i]
}

func (s stringEval) LongestPrefix() string {
	var sequencesGotten string
	var sequence string

	smallerWordIdx, loopSize := s.LoopSize()
	for i := 0; i < loopSize; i++ {
		if s.IndexIsValid(i) {
			sequence += string(s.wordList[smallerWordIdx][i])
			continue
		} 
		if len(sequence) > len(sequencesGotten) {
			sequencesGotten = sequence
		}
		sequence = ""
	}
	if len(sequence) > len(sequencesGotten) {
		sequencesGotten = sequence
	}
	return sequencesGotten
}
