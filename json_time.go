package jsondt

import (
	"fmt"
	"time"
)

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format("15:04:05")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	if len(b) != 10 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("types: failed to unmarshal non-string value %q as an hh:mm:ss", b)
	}
	tm, err := time.Parse("15:04:05", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Time(tm)
	return nil
}
