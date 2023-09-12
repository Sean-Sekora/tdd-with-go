package wordz

var MaximumNumberAllowedGuesses = 5

type Game struct {
	player        Player
	targetWord    Word
	attemptNumber int
	isGameOver    bool
}

type GameOverError struct {
}

func (g *GameOverError) Error() string {
	return "game over"
}

func (g *Game) Attempt(guess string) (Score, error) {
	if g.IsGameOver() {
		return Score{}, &GameOverError{}
	}
	g.trackNumberOfAttempts()

	word := g.targetWord
	score := word.Guess(guess)
	if score.allCorrect() {
		g.endGame()
	} else if g.IsGameOver() {
		return Score{}, &GameOverError{}
	}

	return score, nil
}

func (g *Game) GetAttemptNumber() int {
	return g.attemptNumber
}

func (g *Game) IsGameOver() bool {
	if g.attemptNumber >= MaximumNumberAllowedGuesses {
		g.endGame()
	}
	return g.isGameOver
}

func (g *Game) trackNumberOfAttempts() {
	g.IsGameOver()
	g.attemptNumber++
}

func (g *Game) endGame() {
	g.isGameOver = true
}
