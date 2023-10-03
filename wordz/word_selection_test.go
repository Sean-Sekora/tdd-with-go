package wordz

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

var HIGHEST_WORD_NUMBER = 1
var WORD_NUMBER_SHINE = 2

var db *sql.DB

// setup instantiates a Postgres docker container and attempts to connect to it via a new adapter
func setup() *sql.DB {
	container, err := StartPostgres()
	if err != nil {
		log.Fatalf("Could not start postgres: %s", err)
	}
	defer container.Terminate(context.Background())

	endpoint, err := container.Endpoint(context.Background(), "")
	port := strings.Split(endpoint, ":")[1]

	db, err := ConnectToDB(port)
	if err != nil {
		log.Fatalf("Could not connect to postgres: %s", err)
	}
	defer db.Close()

	// Verify that connection to postgre is successfully established
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping postgres: %s", err)
	}

	return db
}

func TestMain(m *testing.M) {
	db = setup() // Setup one container for test suite to limit resources created during tests
	code := m.Run()
	os.Exit(code)
}

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

func TestWordSelection_FetchWord(t *testing.T) {
	//repository := WordRepositoryPostgres{db: db}
	actual := "" //repository.FetchWordByNumber(1)

	assert.Equal(t, "ARISE", actual)
}
