package wordz_test

import (
	"reflect"
	"tdd_with_go/wordz"
	"testing"
)

func TestWord_guess(t *testing.T) {
	type testCase struct {
		correctWord string
		attempt     string
		expected    []wordz.Letter
	}

	tests := map[string]testCase{
		"incorrect guess": {
			correctWord: "Cat",
			attempt:     "Dog",
			expected: []wordz.Letter{
				wordz.INCORRECT, wordz.INCORRECT, wordz.INCORRECT,
			},
		},
		"correct guess": {
			correctWord: "Cat",
			attempt:     "Cat",
			expected: []wordz.Letter{
				wordz.CORRECT, wordz.CORRECT, wordz.CORRECT,
			},
		},
		"part correct": {
			correctWord: "Cat",
			attempt:     "taC",
			expected: []wordz.Letter{
				wordz.PART_CORRECT, wordz.CORRECT, wordz.PART_CORRECT,
			},
		},
		"rune test": {
			correctWord: "日本語",
			attempt:     "日本語",
			expected: []wordz.Letter{
				wordz.CORRECT, wordz.CORRECT, wordz.CORRECT,
			},
		},
	}

	for name, rtc := range tests {
		tc := rtc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			word := wordz.Word(tc.correctWord)
			score := word.Guess(tc.attempt)

			if assertScore(score, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, score.Results)
			}
		})
	}
}

func assertScore(score wordz.Score, expected []wordz.Letter) bool {
	return !reflect.DeepEqual(score.Results, expected)
}
