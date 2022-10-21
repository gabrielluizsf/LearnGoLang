package main

import (
	"fmt"
	"math"
)

func main(){
 Calcular_Menssagens()
 printBinary(42)
 printHexadecimal(42)
 printHexadecimal(31337)
 fmt.Println("[LVL2]: THE END")
}

func printBinary(value int){
	fmt.Printf("Número: %v\tBinário: %b\n",value,value)
}
func printHexadecimal(value int){
	fmt.Printf("Número: %v\tHexadecimal: %#x\n",value,value)
}

func Calcular_Menssagens(){
fmt.Printf("1 Lampada:")
	numberOfLamps(1)
fmt.Printf("2 Lampadas:")
	numberOfLamps(2)
fmt.Printf("3 Lampadas:")
	numberOfLamps(3)
fmt.Printf("4 Lampadas:")
	numberOfLamps(4)
fmt.Printf("5 Lampadas:")
	numberOfLamps(5)
fmt.Printf("6 Lampadas:")
	numberOfLamps(6)
fmt.Printf("7 Lampadas:")
	numberOfLamps(7)
fmt.Printf("8 Lampadas:")
	numberOfLamps(8)
} 
func numberOfLamps(lampadas float64){
	messages := math.Pow(2,lampadas)
fmt.Printf("\nmensagens:")
	fmt.Println(messages)
fmt.Printf("\n")
}