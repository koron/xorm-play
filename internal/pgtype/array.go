package pgtype

import (
	"fmt"

	"github.com/lib/pq"
)

type StringArray pq.StringArray

func (sa *StringArray) FromDB(b []byte) error {
	return (*pq.StringArray)(sa).Scan(b)
}

func (sa *StringArray) ToDB() ([]byte, error) {
	v, err := (*pq.StringArray)(sa).Value()
	if err != nil {
		return nil, err
	}
	switch w := v.(type) {
	case []byte:
		return w, nil
	default:
		return nil, fmt.Errorf("unsupported type for StringArray: %T", v)
	}
}

func (sa *StringArray) Append(s ...string) {
	*sa = append(*sa, s...)
}
