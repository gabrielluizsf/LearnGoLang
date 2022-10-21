package main

import (
	"fmt"
)

type inteiro int

var(
	numero inteiro  = 1000
	numeroEquivalente int
)

func main(){
	numeroEquivalente = 10;
	fmt.Printf("Valor: %v Tipo: %T\n",numeroEquivalente,numeroEquivalente);
	fmt.Printf("Valor: %v Tipo: %T\n",numero,numero);
	numeroEquivalente = int(numero);
	fmt.Printf("Valor: %v Tipo: %T\n",numeroEquivalente,numeroEquivalente);
	result := numero == 1000;
	fmt.Printf("True/False: %v Tipo: %T",result,result);
}