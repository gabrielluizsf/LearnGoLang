package main

import (
	"fmt"
)
var(
	caractere string
	booleano bool
	inteiro int
	real float64
)

func main(){
	fmt.Printf("Value: %v\nType:%T\n\n",inteiro,inteiro)
	fmt.Printf("Value: %v\nType:%T\n\n",booleano,booleano)
	fmt.Printf("Value: %v\nType:%T\n\n",caractere,caractere)
	fmt.Printf("Value: %v\nType:%T",real,real)

}