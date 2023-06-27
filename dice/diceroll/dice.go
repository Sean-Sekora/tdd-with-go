package diceroll

type Dice struct {
	numberOfSides int
	randomNumber  RandomNumbers
}

func NewDice(numberOfSides int, randomNumber RandomNumbers) *Dice {
	return &Dice{numberOfSides, randomNumber}
}

func (dice Dice) Roll() int {
	return dice.randomNumber.NextInt(dice.numberOfSides)
}
