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

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS bar (
		id  BIGSERIAL PRIMARY KEY,
		geo GEOGRAPHY(POINT, 4326)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS "IDX_bar_geo"
		ON bar USING GiST (geo)`)
	if err != nil {
		log.Fatal(err)
	}
}
