package wordz

import "database/sql"

type WordRepositoryPostgres struct {
	db *sql.DB
}

func (r *WordRepositoryPostgres) FetchWordByNumber(i int) string {
	return ""
}

func (r *WordRepositoryPostgres) HighestWordNumber() int {
	return 0
}

func (r *WordRepositoryPostgres) InsertWord(word string) {
}
