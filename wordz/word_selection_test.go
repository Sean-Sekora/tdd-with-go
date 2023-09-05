package wordz

import "testing"

var HIGHEST_WORD_NUMBER = 1
var WORD_NUMBER_SHINE = 2

func TestWordSelection_ChooseRandomWord(t *testing.T) {
	mockedWordRepository := NewWordRepositoryMock()
	mockedWordRepository.On("FetchWordByNumber", WORD_NUMBER_SHINE).Return("SHINE")
	mockedWordRepository.On("HighestWordNumber").Return(HIGHEST_WORD_NUMBER)

	type fields struct {
		Repository WordRepository
		Random     RandomNumbers
	}
	tests := []struct {
		name   string
		fields fields
		want   Word
	}{
		{
			name: "Happy Path",
			fields: fields{
				Repository: mockedWordRepository,
				Random:     RandomNumbersStub{WORD_NUMBER_SHINE},
			},
			want: Word("SHINE"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WordSelection{
				Repository: tt.fields.Repository,
				Random:     tt.fields.Random,
			}
			if got := ws.ChooseRandomWord(); got != tt.want {
				t.Errorf("ChooseRandomWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
