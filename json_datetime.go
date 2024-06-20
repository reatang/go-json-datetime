package jsondt

import (
	"fmt"
	"time"
)

type DateTime time.Time

func (t DateTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	if len(b) != 21 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("types: failed to unmarshal non-string value %q as an DateTime", b)
	}
	tm, err := time.Parse("2006-01-02 15:04:05", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = DateTime(tm)
	return nil
}
