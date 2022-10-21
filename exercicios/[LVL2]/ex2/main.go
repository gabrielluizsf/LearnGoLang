package main

import (
	"fmt"
)
const(
	dez = 10
	vinte = 20
)
var (
	igual =  dez == vinte
	diferente =  dez != vinte
	maior = dez >  vinte
	maiorOuIgual =  dez >= vinte
	menor =	 dez <  vinte
	menorOuIgual =	dez <= vinte
)
func main(){
	fmt.Printf("É igual: %v\n",igual)
	fmt.Printf("É diferente: %v\n",diferente)
	fmt.Printf("É maior: %v\n",maior)
	fmt.Printf("É menor: %v\n",menor)
	fmt.Printf("É maior ou igual: %v\n",maiorOuIgual)
	fmt.Printf("É menor ou igual: %v\n",menorOuIgual)
}