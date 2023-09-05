package wordz

import "github.com/stretchr/testify/mock"

type GameRepositoryMock struct {
	mock.Mock
}

func NewGameRepositoryMock() *GameRepositoryMock {
	return &GameRepositoryMock{}
}

func (r *GameRepositoryMock) Create(game Game) Game {
	ret := r.Called(game)

	var r0 Game
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(Game)
	}

	return r0
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
