package main

import (
	"tdd_with_go/dice/diceroll"
	"tdd_with_go/zapadapter"
)

var logger = zapadapter.NewZapAdapter()

func main() {
	rnd := diceroll.RandomNumbersImpl{}
	dice := diceroll.NewDice(6, rnd)
	roll := dice.Roll()
	logger.Info("Rolling a dice", "roll", roll)
}
