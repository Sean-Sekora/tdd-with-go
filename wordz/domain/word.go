package domain

type Word struct {
	word string
}

func newWord(correctWord string) Word {
	return Word{word: correctWord}
}

func (word Word) guess(attempt string) Score {
	score := newScore(word.word)
	score.assess(attempt)

	return score
}
