package bean

import "github.com/koron/xorm-play/internal/pgtype"

type Bar struct {
	ID  int64            `xorm:"id BIGSERIAL PK AUTOINCR"`
	Geo *pgtype.GeoPoint `xorm:"geo BLOB"`
}
