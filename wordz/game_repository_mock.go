package wordz

import "github.com/stretchr/testify/mock"

type GameRepositoryMock struct {
	mock.Mock
}

func NewGameRepositoryMock() *GameRepositoryMock {
	return &GameRepositoryMock{}
}

func (r *GameRepositoryMock) Create(player Player, word Word) Game {
	return Game{
		player:        player,
		targetWord:    word,
		attemptNumber: 0,
		isGameOver:    false,
	}
}

func (r *GameRepositoryMock) FetchForPlayer(player Player) Game {
	ret := r.Called(player)

	var r0 Game
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(Game)
	}

	return r0
}

func (r *GameRepositoryMock) Update(game Game) {
	r.Called(game)
}
