package jsondt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTime_String(t1 *testing.T) {
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
				Time: time.Date(0, 0, 0, 12, 12, 12, 0, time.UTC),
			},
			want: `12:12:12`,
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

type TestJSONTime struct {
	Time Time `json:"time"`
}

func TestTime_MarshalJSON(t1 *testing.T) {
	data := TestJSONTime{
		Time: Time{time.Date(0, 0, 0, 1, 2, 3, 0, time.UTC)},
	}

	bytes, err := json.Marshal(data)
	fmt.Println(string(bytes), err)
}

func TestTime_UnmarshalJSON(t1 *testing.T) {
	j := []byte(`{"time":"01:02:03"}`)
	var data TestJSONTime

	err := json.Unmarshal(j, &data)
	fmt.Println(data.Time.Hour(), data.Time.Minute(), data.Time.Second(), err)
}
