package main

import (
	"fmt"
)

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB 
	GB 
	TB  
)

func main(){

	fmt.Println("Deslocando os Bits:")
	fmt.Printf("KB em Binário: %b\t\t\t\t\tKB em Decimal: %d\n",KB,KB)
	fmt.Printf("MB em Binário: %b\t\t\t\tMB em Decimal: %d\n",MB,MB)
	fmt.Printf("GB em Binário: %b\t\t\tGB em Decimal: %d\n",GB,GB)
	fmt.Printf("TB em Binário  %b\tTB em Decimal: %d\n",TB,TB)
}