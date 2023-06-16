package domain

import "strings"

type Score struct {
	CorrectWord Word
	Results     []Letter
	Position    int
}

func newScore(correctWord Word) Score {
	return Score{
		CorrectWord: correctWord,
		Results:     make([]Letter, len(correctWord)),
	}
}

func (s Score) assess(attempt string) {
	for i, character := range attempt {
		s.Position = i
		s.Results[s.Position] = s.scoreFor(character)
	}
}

func (s Score) scoreFor(character rune) Letter {
	if s.isCorrect(character) {
		return CORRECT
	} else if s.occursInWord(character) {
		return PART_CORRECT
	} else {
		return INCORRECT
	}
}

func (s Score) isCorrect(attempt rune) bool {
	return s.CorrectWord[s.Position] == byte(attempt)
}

func (s Score) occursInWord(attempt rune) bool {
	return strings.ContainsRune(string(s.CorrectWord), attempt)
}
