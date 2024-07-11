package jsondt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type testJSONDateTime struct {
	CreatedAt DateTime `json:"created_at"`
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	data := testJSONDateTime{
		CreatedAt: DateTime{time.Date(2020, time.April, 28, 1, 2, 3, 0, time.UTC)},
	}

	bytes, err := json.Marshal(data)
	fmt.Println(string(bytes), err)
}

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	j := []byte(`{"created_at":"2020-04-28 01:02:03"}`)
	var data testJSONDateTime

	err := json.Unmarshal(j, &data)
	fmt.Println(data.CreatedAt.String(), err)
}

type testJSONDateTime2 struct {
	CreatedAt *DateTime `json:"created_at"`
}

func TestJSONTime_UnmarshalJSON2(t *testing.T) {
	data := testJSONDateTime2{
		CreatedAt: nil,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	if string(bytes) != `{"created_at":null}` {
		t.Error("error: ", string(bytes))
	}
}

func TestJSONTime_UnmarshalJSONNull(t *testing.T) {
	j := []byte(`{"created_at":null}`)
	var data testJSONDateTime2

	err := json.Unmarshal(j, &data)
	fmt.Println(data.CreatedAt == nil, err)
}
