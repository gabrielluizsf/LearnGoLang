package main

import (
	operadoreslogicos "condicional-golang/operadores-logicos"
	"fmt"
)

func main(){
	isTrue();
    Loading(10);
    Cases("apple");
    operadoreslogicos.LogicOperators(6)
}
func isTrue(){
	booleano := true
	if booleano == true{
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
}
//! == NOT
func Loading(value int){
  if !(value > 100){
      fmt.Println("Loading....")
  }else if value > 100{
      fmt.Println("Invalid Value")
  }else{
      fmt.Println("100%")
    }
}

func Cases(choice string){

    switch choice{
    case "bread":
        fmt.Println("Bread: R$2.00");
    case "apple":
        fmt.Println("Apple: R$10.00")
        fallthrough
    case "limon":
        fmt.Println("Limon: FREE");
    case "foods":
        fmt.Println("Foods: bread | apple")
    default:
        fmt.Println("Invalid Choice")
    }

}

