package main

import (
	"encoding/json"
	"os"
)

type animal struct{
	Name  string
	Color string
}

func main(){
	gopher := animal{
		Name: "Gopher",
		Color: "Blue",	
	}
	encoder := json.NewEncoder(os.Stdout);
	encoder.Encode(gopher)
}