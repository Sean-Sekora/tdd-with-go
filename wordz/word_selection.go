package wordz

type WordSelection struct {
	Repository WordRepository
	Random     RandomNumbers
}

func (ws WordSelection) ChooseRandomWord() Word {
	wordNumber := ws.Random.NextInt(ws.Repository.HighestWordNumber())
	return Word(ws.Repository.FetchWordByNumber(wordNumber))
}
