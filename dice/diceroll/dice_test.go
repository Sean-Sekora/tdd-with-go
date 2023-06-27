package diceroll_test

import (
	"tdd_with_go/dice/diceroll"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDice_Roll(t *testing.T) {
	type fields struct {
		numberOfSides int
		randomNumber  diceroll.RandomNumbers
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Roll a 6-sided dice",
			fields: fields{
				numberOfSides: 6,
				randomNumber:  &diceroll.RandomNumbersStub{},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dice := diceroll.NewDice(tt.fields.numberOfSides, tt.fields.randomNumber)
			assert.Equal(t, tt.want, dice.Roll())
		})
	}
}
