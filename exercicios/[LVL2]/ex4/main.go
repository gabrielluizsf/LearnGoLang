package main

import (
	"fmt"
)
var (
	valor int = 42
 	valorDeslocado = valor << 1
)

func main(){
	fmt.Printf("Decimal: %d\tBinário: %b\tHexadecimal: %#x\n",valor,valor,valor)
	fmt.Printf("Decimal: %d\tBinário %b\tHexadecimal: %#x\n",valorDeslocado,valorDeslocado,valorDeslocado)

}
