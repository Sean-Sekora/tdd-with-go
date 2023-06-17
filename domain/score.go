package domain

import (
	"strings"
)

type Score struct {
	CorrectWord Word
	Results     []Letter
	Position    int
}

func (s *Score) assess(attempt string) {
	for i, character := range attempt {
		s.Position = i
		s.Results = append(s.Results, s.scoreFor(character))
	}
}

func (s *Score) scoreFor(character rune) Letter {
	switch {
	case s.isCorrect(character):
		return CORRECT
	case s.occursInWord(character):
		return PART_CORRECT
	default:
		return INCORRECT
	}
}

func (s *Score) isCorrect(attempt rune) bool {
	return strings.IndexRune(string(s.CorrectWord), attempt) == s.Position
}

func (s *Score) occursInWord(attempt rune) bool {
	return strings.ContainsRune(string(s.CorrectWord), attempt)
}
