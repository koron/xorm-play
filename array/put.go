package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/koron/xorm-play/internal/bean"
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

	vs := flag.Args()
	if len(vs) == 0 {
		vs = append(vs, "sample", time.Now().Format(time.RFC3339))
	}

	n, err := eng.Insert(bean.NewFoo(vs...))
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println(n)
}
