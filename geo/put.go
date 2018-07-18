package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/koron/xorm-play/internal/bean"
	"github.com/koron/xorm-play/internal/pgtype"
	"github.com/kr/pretty"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()

	eng, err := xorm.NewEngine("postgres", "postgres://postgres@127.0.0.1:8000/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer eng.Close()

	latlng := flag.Args()
	if len(latlng) < 2 {
		log.Fatal("need two more args")
	}
	lat, err := strconv.ParseFloat(latlng[0], 64)
	if err != nil {
		log.Fatal(err)
	}
	lng, err := strconv.ParseFloat(latlng[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	n, err := eng.Insert(&bean.Bar{Geo: pgtype.NewGeoPoint(lat, lng)})
	if err != nil {
		log.Fatal("insert failed: ", err)
	}
	pretty.Println(n)
}
