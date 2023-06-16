package domain

import "strings"

type Score struct {
	correctWord string
	results     []Letter
	position    int
}

func newScore(correctWord string) Score {
	return Score{
		correctWord: correctWord,
		results:     make([]Letter, len(correctWord)),
	}
}

func (s Score) assess(attempt string) {
	for i, character := range attempt {
		s.position = i
		s.results[s.position] = s.scoreFor(character)
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
	return s.correctWord[s.position] == byte(attempt)
}

func (s Score) occursInWord(attempt rune) bool {
	return strings.ContainsRune(s.correctWord, attempt)
}
