package wordz

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWordz_NewGame(t *testing.T) {
	// Arrange
	player := Player{}
	game := Game{
		player:        player,
		targetWord:    "ARISE",
		attemptNumber: 0,
		isGameOver:    false,
	}
	mockedGameRepository := NewGameRepositoryMock()
	mockedGameRepository.On("Create", mock.Anything).Return(game)
	mockedWordRepository := NewWordRepositoryMock()
	mockedWordRepository.On("FetchWordByNumber", mock.Anything).Return("ARISE")
	mockedWordRepository.On("HighestWordNumber").Return(1)

	wordSelection := WordSelection{
		Repository: mockedWordRepository,
		Random:     RandomNumbersStub{1},
	}
	wordz := Wordz{
		gameRepository: mockedGameRepository,
		wordSelection:  wordSelection,
	}

	// Act
	newGame := wordz.NewGame(player)

	//Assert
	assert.Equal(t, game, newGame)
}
