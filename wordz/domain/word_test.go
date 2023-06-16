package domain

import (
	"reflect"
	"testing"
)

func TestWord_guess(t *testing.T) {

	type testCase struct {
		correctWord string
		attempt     string
		expected    []Letter
	}

	tests := map[string]testCase{
		"incorrect guess": {
			correctWord: "cat",
			attempt:     "dog",
			expected:    []Letter{INCORRECT, INCORRECT, INCORRECT},
		},
		"correct guess": {
			correctWord: "cat",
			attempt:     "cat",
			expected:    []Letter{CORRECT, CORRECT, CORRECT},
		},
		"part correct": {
			correctWord: "cat",
			attempt:     "tac",
			expected:    []Letter{PART_CORRECT, CORRECT, PART_CORRECT},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			word := newWord(test.correctWord)
			score := word.guess(test.attempt)
			if assertScore(score, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, score.results)
			}
		})
	}
}

func assertScore(score Score, expected []Letter) bool {
	return !reflect.DeepEqual(score.results, expected)
}
