package main

import (
	"log"

	"github.com/go-xorm/xorm"
	"github.com/koron/xorm-play/internal/bean"
	"github.com/kr/pretty"
	_ "github.com/lib/pq"
)

func main() {
	eng, err := xorm.NewEngine("postgres", "postgres://postgres@127.0.0.1:8000/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer eng.Close()

	var list []*bean.Foo
	err = eng.Find(&list)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println(list)
}
