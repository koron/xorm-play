package pgtype

import (
	"encoding/hex"
	"fmt"

	postgis "github.com/cridenour/go-postgis"
)

type GeoPoint postgis.PointS

func NewGeoPoint(lat, lng float64) *GeoPoint {
	return &GeoPoint{
		SRID: 4326,
		X:    lat,
		Y:    lng,
	}
}

func (gp *GeoPoint) FromDB(b []byte) error {
	return (*postgis.PointS)(gp).Scan(b)
}

func (gp *GeoPoint) ToDB() ([]byte, error) {
	v, err := (*postgis.PointS)(gp).Value()
	if err != nil {
		return nil, err
	}
	switch w := v.(type) {
	case []byte:
		dst := make([]byte, hex.EncodedLen(len(w)))
		hex.Encode(dst, w)
		return dst, nil
	default:
		return nil, fmt.Errorf("unsupported type for GeoPoint: %T", v)
	}
}
