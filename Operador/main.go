package main;

import (
	"fmt"
)

var( 
	i = 10
	y = true
)

func main(){
	x := 10.1
	z := "z"
	e := 10 == 50


	//Valores
	fmt.Println("VALUES")
	fmt.Printf("E: %v\n", e)
	fmt.Printf("I: %v\n", i)
	fmt.Printf("X: %v\n", x)
	fmt.Printf("Y: %v\n", y)
	fmt.Printf("Z: %v\n\n", z)

	//Tipos
	fmt.Println("TYPES")
	fmt.Printf("E: %T\n", e)
	fmt.Printf("I: %T\n", i)
	fmt.Printf("X: %T\n", x)
	fmt.Printf("Y: %T\n", y)
	fmt.Printf("Z: %T\n", z)
}