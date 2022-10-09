package main

import (
	"fmt"
)

type inteiro int

var (
	y int
	x inteiro 
)

func main(){
	fmt.Printf("----------------------\n");
	fmt.Printf("Valor de X: %v \nTipo de X: %T\n",x,x);
	x = 42;
	fmt.Printf("Valor de X: %v\n",x);
	fmt.Printf("----------------------\n");
	y = int(x);
	fmt.Printf("Valor de Y: %v \nTipo de Y: %T\n",y,y);
	fmt.Printf("-----------------------\n");
}