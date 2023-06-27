package diceroll

type RandomNumbersStub struct {
}

func (RandomNumbersStub) NextInt(int) int {
	return 4
}
