package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

const apiURL = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

// Timestamp _
type Timestamp time.Time

// UnmarshalJSON _
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func main() {
	val := make(map[string]Timestamp)
	if err := json.Unmarshal([]byte(apiURL), &val); err != nil {
		panic(err)
	}
	fmt.Println(val)
	for k, v := range val {
		fmt.Printf("%s ----type %v \n", k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
}
