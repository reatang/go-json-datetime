package jsondt

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ json.Marshaler = (*Time)(nil)
var _ json.Unmarshaler = (*Time)(nil)

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format("15:04:05")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if isEmpty(b) {
		return nil
	}

	if len(b) != 10 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("%w: failed to unmarshal non-string value %q as an hh:mm:ss", ErrJSONDateTime, b)
	}

	tm, err := time.Parse("15:04:05", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Time(tm)
	return nil
}
