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
		targetWord:    "BINGO",
		attemptNumber: 0,
		isGameOver:    false,
	}
	mockedGameRepository := NewGameRepositoryMock()
	mockedWordRepository := NewWordRepositoryMock()
	mockedWordRepository.On("FetchWordByNumber", 1).Return("Not Me")
	mockedWordRepository.On("FetchWordByNumber", 2).Return("BINGO")
	mockedWordRepository.On("HighestWordNumber").Return(2)

	wordSelection := WordSelection{
		Repository: mockedWordRepository,
		Random:     RandomNumbersStub{2},
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

func TestWordz_Guess(t *testing.T) {
	type fields struct {
		gameRepository GameRepository
		wordSelection  WordSelection
	}
	type args struct {
		player Player
		guess  string
	}

	mockedGameRepository := NewGameRepositoryMock()
	mockedGameRepository.On("FetchForPlayer", mock.Anything).Return(Game{
		player:        Player{},
		targetWord:    "BINGO",
		attemptNumber: 0,
		isGameOver:    false,
	})
	mockedGameRepository.On("Update", mock.Anything).Return()

	tests := []struct {
		name   string
		fields fields
		args   args
		want   GuessResult
		want1  GameOverError
	}{
		{
			name: "Happy Path",
			fields: fields{
				gameRepository: mockedGameRepository,
				wordSelection:  WordSelection{},
			},
			args: args{
				player: Player{},
				guess:  "BINGO",
			},
			want: GuessResult{
				Score: Score{
					CorrectWord: "BINGO",
					Results:     []Letter{CORRECT, CORRECT, CORRECT, CORRECT, CORRECT},
					Position:    4,
				},
				IsGameOver: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wordz{
				gameRepository: tt.fields.gameRepository,
				wordSelection:  tt.fields.wordSelection,
			}
			got, got1 := w.Guess(tt.args.player, tt.args.guess)
			assert.Equalf(t, tt.want, got, "Guess(%v, %v)", tt.args.player, tt.args.guess)
			assert.Equalf(t, tt.want1, got1, "Guess(%v, %v)", tt.args.player, tt.args.guess)
		})
	}
}
