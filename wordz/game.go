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

func (g GameOverError) Error() string {
	panic("Game Over Error")
}

func (g *Game) Attempt(guess string) (Score, error) {
	if g.isGameOver {
		return Score{}, GameOverError{}
	}
	g.trackNumberOfAttempts()
	word := g.targetWord
	score := word.Guess(guess)
	if score.allCorrect() {
		g.endGame()
	}

	return score, nil
}

func (g *Game) GetAttemptNumber() int {
	return g.attemptNumber
}

func (g *Game) IsGameOver() bool {
	return g.isGameOver
}

func (g *Game) trackNumberOfAttempts() {
	g.attemptNumber++

	if g.attemptNumber >= MaximumNumberAllowedGuesses {
		g.endGame()
	}
}

func (g *Game) endGame() {
	g.isGameOver = true
}
