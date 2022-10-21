package main;

import (
	"fmt"
);

type inteiro int;

var(
	x inteiro;
);

func main(){
	fmt.Printf("Valor de X: %v \nTipo: %T\n",x,x);
	x = 42;
	fmt.Printf("Valor de X: %v",x);
}
