package jsondt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDate_String(t1 *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				Time: time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
			},
			want: `2024-12-12`,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Time{
				Time: tt.fields.Time,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestJSONDate struct {
	Date Date `json:"date"`
}

func TestDate_MarshalJSON(t1 *testing.T) {
	data := TestJSONDate{
		Date: Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)},
	}

	bytes, err := json.Marshal(data)
	fmt.Println(string(bytes), err)
}

func TestDate_UnmarshalJSON(t1 *testing.T) {
	j := []byte(`{"date":"2024-12-12"}`)
	var data TestJSONDate

	err := json.Unmarshal(j, &data)
	fmt.Println(data.Date.Year(), int(data.Date.Month()), data.Date.Day(), err)
}
