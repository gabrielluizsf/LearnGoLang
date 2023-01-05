package main

import (
  "fmt"
  "encoding/json"
)



type person struct {
	First   string
	Last    string
	Sayings []string
}

var default_error = "[Warning]: ";

func main() {
	people := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	};

	bs, err := toJSON(people);
    if err != nil{
    fmt.Errorf(default_error, err);
  }
	fmt.Println(string(bs));

}

func toJSON(a interface{}) ([]uint8, error) {
	bs, err := json.Marshal(a);
	if err != nil {
    return []byte{}, fmt.Errorf(default_error, err);
	}
	return bs, nil;
}
