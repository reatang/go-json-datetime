package jsondt

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ json.Marshaler = (*Date)(nil)
var _ json.Unmarshaler = (*Date)(nil)

type Date time.Time

func (t Date) String() string {
	return time.Time(t).Format("2006-01-02")
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Date) UnmarshalJSON(b []byte) error {
	if isNull(b) {
		return nil
	}

	if len(b) != 12 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("%w: failed to unmarshal non-string value %q as an YYYY-MM-dd", ErrJSONDateTime, b)
	}

	tm, err := time.Parse("2006-01-02", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Date(tm)
	return nil
}
