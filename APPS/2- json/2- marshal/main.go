package main

import( 
	"fmt"
	"encoding/json"
)
type People struct{
	Name string
	Age  int
	Job  string
	Bank_Account float64

}

func main(){
	james_bond := People{
		Name: "James Bond",
		Age:   40,
		Job: "Agente Secreto",
		Bank_Account: 1000700,
	}

	darth_vader := People{
		Name: "Anakin Skywalker",
		Age:   50,
		Job: "Comandante Supremo das Forças Imperiais",
		Bank_Account: 57497589472572567,
	}
	james, err := json.Marshal(james_bond);
		printError(err);
	anakin, err := json.Marshal(darth_vader);
		printError(err);
	fmt.Println(string(james));
	fmt.Println(string(anakin));
}
func printError(err error){
	if err != nil{
		fmt.Println(err)
	}
}