package domain

type Word struct {
	word string
}

func (word Word) guess(attempt string) Score {
	score := newScore(word.word)
	score.assess(attempt)

	return score
}
