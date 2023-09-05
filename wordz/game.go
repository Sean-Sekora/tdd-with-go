package wordz

var MaximumNumberAllowedGuesses = 5

type Game struct {
	player        Player
	targetWord    Word
	attemptNumber int
	isGameOver    bool
}

func (g *Game) Attempt(guess string) Score {
	g.trackNumberOfAttempts()
	word := g.targetWord
	score := word.Guess(guess)
	if score.allCorrect() {
		g.endGame()
	}

	return score
}

func (g *Game) trackNumberOfAttempts() {
	g.attemptNumber++

	if g.attemptNumber > MaximumNumberAllowedGuesses {
		g.endGame()
	}
}

func (g *Game) endGame() {
	g.isGameOver = true
}
