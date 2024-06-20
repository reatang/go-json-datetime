package jsondt

import (
	"fmt"
	"time"
)

type Date time.Time

func (t Date) String() string {
	return time.Time(t).Format("2006-01-02")
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	if len(b) != 12 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("types: failed to unmarshal non-string value %q as an YYYY-MM-dd", b)
	}
	tm, err := time.Parse("2006-01-02", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Date(tm)
	return nil
}
