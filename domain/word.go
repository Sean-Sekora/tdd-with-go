package domain

type Word string

func (word Word) Guess(attempt string) Score {
	score := Score{CorrectWord: word}
	score.assess(attempt)

	return score
}
