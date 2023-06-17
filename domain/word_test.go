package domain_test

import (
	"reflect"
	"tdd-with-go/domain"
	"testing"
)

func TestWord_guess(t *testing.T) {
	type testCase struct {
		correctWord string
		attempt     string
		expected    []domain.Letter
	}

	tests := map[string]testCase{
		"incorrect guess": {
			correctWord: "Cat",
			attempt:     "Dog",
			expected: []domain.Letter{
				domain.INCORRECT, domain.INCORRECT, domain.INCORRECT,
			},
		},
		"correct guess": {
			correctWord: "Cat",
			attempt:     "Cat",
			expected: []domain.Letter{
				domain.CORRECT, domain.CORRECT, domain.CORRECT,
			},
		},
		"part correct": {
			correctWord: "Cat",
			attempt:     "taC",
			expected: []domain.Letter{
				domain.PART_CORRECT, domain.CORRECT, domain.PART_CORRECT,
			},
		},
		"rune test": {
			correctWord: "日本語",
			attempt:     "日本語",
			expected: []domain.Letter{
				domain.CORRECT, domain.CORRECT, domain.CORRECT,
			},
		},
	}

	for name, rtc := range tests {
		tc := rtc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			word := domain.Word(tc.correctWord)
			score := word.Guess(tc.attempt)

			if assertScore(score, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, score.Results)
			}
		})
	}
}

func assertScore(score domain.Score, expected []domain.Letter) bool {
	return !reflect.DeepEqual(score.Results, expected)
}
