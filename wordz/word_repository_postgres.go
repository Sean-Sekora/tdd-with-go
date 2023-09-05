package wordz

import (
	"database/sql"
	"fmt"
)

type WordRepositoryPostgres struct {
}

const (
	host     = "localhost"
	port     = 5432
	user     = "ssekora"
	password = ""
	dbname   = "wordz"
)

func (r WordRepositoryPostgres) FetchWordByNumber(n int) string {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`select * from "word" where id = $1`, n)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var word string
		err = rows.Scan(&id, &word)
		if err != nil {
			panic(err)
		}
		return word
	}
	return ""
}

func (r WordRepositoryPostgres) HighestWordNumber() int {
	panic("implement me")
}
