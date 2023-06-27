package wordz

import "fmt"

type Word string

func (word Word) Guess(attempt string) Score {
	score := Score{CorrectWord: word}
	fmt.Printf("Original Memory Address: %p\n", &score)
	score.assess(attempt)

	return score
}
