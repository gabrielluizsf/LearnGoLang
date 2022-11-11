package main

import (
	"fmt"
	"encoding/json"
)

type people struct{
	Name string `json: Name`
	Age  int	`json: Age`
	Job  string	`json: Job`
	Bank_Account float64 `json: Bank_Account`

}
func main(){

	info := []byte(`{"Name":"Anakin Skywalker","Age":50,"Job":"Comandante Supremo das Forças Imperiais","Bank_Account":73944592568743765635764385.4385784356344854675}`);
	
	var anakin people
	err := json.Unmarshal(info, &anakin);
	if err != nil{
		fmt.Println("[ERROR]: ",err);
	}
	fmt.Println(anakin.Job);

}