package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres@127.0.0.1:8000/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS foo (
		id BIGSERIAL PRIMARY KEY,
		vs TEXT[]
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
