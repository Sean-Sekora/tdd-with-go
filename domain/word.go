package domain

type Word string

func (word Word) Guess(attempt string) Score {
	score := newScore(word)
	score.assess(attempt)

	return score
}
