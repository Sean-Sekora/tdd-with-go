package wordz

type RandomNumbersStub struct {
	num int
}

func (r RandomNumbersStub) NextInt(int) int {
	return r.num
}
