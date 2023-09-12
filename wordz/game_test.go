package wordz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_Attempt(t *testing.T) {
	type fields struct {
		player        Player
		targetWord    Word
		attemptNumber int
		isGameOver    bool
	}
	type args struct {
		guess string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Letter
		error  error
	}{
		{
			name: "Happy Path",
			fields: fields{
				player:        Player{},
				targetWord:    "BINGO",
				attemptNumber: 0,
				isGameOver:    false,
			},
			args: args{
				guess: "BINGO",
			},
			want: []Letter{CORRECT, CORRECT, CORRECT, CORRECT, CORRECT},
		},
		{
			name: "Partial Match",
			fields: fields{
				player:        Player{},
				targetWord:    "BINGO",
				attemptNumber: 0,
				isGameOver:    false,
			},
			args: args{
				guess: "BIN",
			},
			want: []Letter{CORRECT, CORRECT, CORRECT},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player:        tt.fields.player,
				targetWord:    tt.fields.targetWord,
				attemptNumber: tt.fields.attemptNumber,
				isGameOver:    tt.fields.isGameOver,
			}
			score, err := g.Attempt(tt.args.guess)
			assert.Equal(t, tt.want, score.Results)
			assert.Equal(t, tt.error, err)
		})
	}
}

func TestGame_IsGameOver(t *testing.T) {
	type fields struct {
		player        Player
		targetWord    Word
		attemptNumber int
		isGameOver    bool
	}
	tests := []struct {
		name    string
		fields  fields
		guesses []string
		want    bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"BINGO"},
			want:    true,
		},
		{
			name: "Partial Match 1",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"ARI"},
			want:    false,
		},
		{
			name: "Partial Match 2",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"BINGOE"},
			want:    false,
		},
		{
			name: "Max Guesses",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"A", "AR", "ARI", "ARIS", "BINGO"},
			want:    true,
		},
		{
			name: "Max Wrong Guesses",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"A", "AR", "ARI", "ARIS", "ARISA"},
			want:    true,
		},
		{
			name: "Too Many Guesses",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"A", "AR", "ARI", "ARIS", "ARISA", "BINGO"},
			want:    true,
		},
		{
			name: "Guess After Game Over",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"BING", "BINGO", "WAS HIS NAMEO"},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player:        tt.fields.player,
				targetWord:    tt.fields.targetWord,
				attemptNumber: tt.fields.attemptNumber,
				isGameOver:    tt.fields.isGameOver,
			}
			for _, guess := range tt.guesses {
				g.Attempt(guess)
			}
			assert.Equalf(t, tt.want, g.IsGameOver(), "IsGameOver()")
		})
	}
}

func TestGame_UpdateAttemptNumber(t *testing.T) {
	type fields struct {
		player        Player
		targetWord    Word
		attemptNumber int
		isGameOver    bool
	}
	tests := []struct {
		name    string
		fields  fields
		guesses []string
		want    int
		error   error
	}{
		{
			name: "1 Guess",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"A"},
			want:    1,
		},
		{
			name: "Multiple Guesses",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"A", "AR", "ARI", "ARIS", "BINGO"},
			want:    5,
		},
		{
			name: "Too Many Guesses",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"A", "AR", "ARI", "ARIS", "ARISA", "BINGO"},
			want:    5,
			error:   &GameOverError{},
		},
		{
			name: "Guess After Game Over",
			fields: fields{
				player:     Player{},
				targetWord: "BINGO",
			},
			guesses: []string{"BING", "BINGO", "WAS HIS NAMEO"},
			want:    2,
			error:   &GameOverError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player:        tt.fields.player,
				targetWord:    tt.fields.targetWord,
				attemptNumber: tt.fields.attemptNumber,
				isGameOver:    tt.fields.isGameOver,
			}
			var err error
			for _, guess := range tt.guesses {
				_, err = g.Attempt(guess)
			}
			assert.Equalf(t, tt.want, g.GetAttemptNumber(), "GetAttemptNumber()")
			assert.Equal(t, tt.error, err)
		})
	}
}
