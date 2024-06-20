package jsondt

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ json.Marshaler = (*DateTime)(nil)
var _ json.Unmarshaler = (*DateTime)(nil)

type DateTime time.Time

func (t DateTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	if isNull(b) {
		return nil
	}

	if len(b) != 21 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("%w: failed to unmarshal non-string value %q as an YYYY-MM-dd hh:mm:ss", ErrJSONDateTime, b)
	}

	tm, err := time.Parse("2006-01-02 15:04:05", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = DateTime(tm)
	return nil
}
