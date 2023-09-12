package wordz

type Wordz struct {
	gameRepository GameRepository
	wordSelection  WordSelection
}

func (w *Wordz) NewGame(player Player) Game {
	return w.gameRepository.Create(player, w.wordSelection.ChooseRandomWord())
}

func (w *Wordz) Guess(player Player, guess string) (GuessResult, GameOverError) {
	game := w.gameRepository.FetchForPlayer(player)

	if game.IsGameOver() {
		return GuessResult{}, GameOverError{}
	}
	score, err := game.Attempt(guess)
	if err != nil {
		return GuessResult{}, GameOverError{}
	}

	w.gameRepository.Update(game)

	return GuessResult{
		Score:      score,
		IsGameOver: game.IsGameOver(),
		IsError:    false,
	}, GameOverError{}
}
