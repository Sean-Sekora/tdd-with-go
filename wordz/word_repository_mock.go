package wordz

import "github.com/stretchr/testify/mock"

type WordRepositoryMock struct {
	mock.Mock
}

func NewWordRepositoryMock() *WordRepositoryMock {
	return &WordRepositoryMock{}
}
func (w *WordRepositoryMock) FetchWordByNumber(arg0 int) string {
	ret := w.Called(arg0)

	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	return r0
}

func (w *WordRepositoryMock) HighestWordNumber() int {
	ret := w.Called()

	var r0 int
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(int)
	}

	return r0
}
