package bean

import "github.com/koron/xorm-play/internal/pgtype"

type Foo struct {
	ID int64              `xorm:"id BIGSERIAL PK AUTOINCR"`
	VS pgtype.StringArray `xorm:"vs BLOB"`
}

func NewFoo(s ...string) *Foo {
	v := &Foo{}
	v.VS.Append(s...)
	return v
}
