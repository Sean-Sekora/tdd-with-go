package wordz

type Game struct {
	player        Player
	targetWord    string
	attemptNumber int
	isGameOver    bool
}

func (g *Game) Attempt(guess String) {
	g.isGameOver = false
}

func (g *Game) trackNumberOfAttempts() {
	g.attemptNumber++
}
