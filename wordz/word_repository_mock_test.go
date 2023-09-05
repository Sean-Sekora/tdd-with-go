package wordz

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWordRepositoryMock_FetchWordByNumber(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		arg0 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Happy Path",
			args: args{
				arg0: 1,
			},
			want: "ARISE",
		},
	}
	// Arrange
	mockedWordRepository := NewWordRepositoryMock()
	mockedWordRepository.On("FetchWordByNumber", mock.Anything).Return("ARISE")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mockedWordRepository.FetchWordByNumber(tt.args.arg0); got != tt.want {
				t.Errorf("FetchWordByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
