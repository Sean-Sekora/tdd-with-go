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
			want: 4,
		},
	}
	for _, rt := range tests {
		test := rt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			dice := diceroll.NewDice(test.fields.numberOfSides, test.fields.randomNumber)
			value := dice.Roll()
			assert.Equal(t, test.want, value)
		})
	}
}
