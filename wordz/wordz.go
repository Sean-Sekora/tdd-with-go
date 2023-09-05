package wordz

type Wordz struct {
	gameRepository GameRepository
	wordSelection  WordSelection
}

func (w *Wordz) NewGame(player Player) Game {
	return w.gameRepository.Create(Game{
		player:        player,
		targetWord:    w.wordSelection.ChooseRandomWord(),
		attemptNumber: 0,
		isGameOver:    false,
	})
}
