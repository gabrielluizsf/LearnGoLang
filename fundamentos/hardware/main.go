package main

import (
	"fmt"
	"math"
)

var (
	n float64
	state bool
)
func print(value string){
	fmt.Printf(value);
}
func ligado(){
	state = true
	fmt.Println(state)
}
func desligado(){
	state = false
	fmt.Println(state)
}

func main(){

	print("\n")
	Calcular_Menssagens()
	printarLetra([3]int{0,1,0});
	printarLetra([3]int{1,1,1});
	printarLetra([3]int{1,1,0});
	printarLetra([3]int{0,1,1});
	printarLetra([3]int{0,0,0});
	printarLetra([3]int{1,0,0});
	printarLetra([3]int{1,1,0});
	printarLetra([3]int{0,1,0});
	printarLetra([3]int{0,0,0});
	printarLetra([3]int{1,0,1});
	printarLetra([3]int{1,1,0});

}
func Calcular_Menssagens(){
	print("1 Lampada:")
	numberOfLamps(1)
	print("2 Lampadas:")
	numberOfLamps(2)
	print("3 Lampadas:")
	numberOfLamps(3)
	print("4 Lampadas:")
	numberOfLamps(4)
	print("5 Lampadas:")
	numberOfLamps(5)
	print("6 Lampadas:")
	numberOfLamps(6)
	print("7 Lampadas:")
	numberOfLamps(7)
	print("8 Lampadas:")
	numberOfLamps(8)
}
func numberOfLamps(lampadas float64){
	messages := math.Pow(2,lampadas)
	print("\nmensagens:")
	fmt.Println(messages)
	print("\n")
}
func Checar_Padrao_Printando_A_Frase(pattern [2]bool){

}
func printarLetra(pattern [3]int){
	A := [3]int{0,0,0}
	B := [3]int{0,0,1}
	C := [3]int{0,1,0}
	D := [3]int{1,0,0}
	E := [3]int{1,1,0}
	F := [3]int{1,0,1}
	G := [3]int{0,1,1}
	H := [3]int{1,1,1}


	switch pattern {
	case A:
		fmt.Println("A")
	case B:
		fmt.Println("B")
	case C:
		fmt.Println("C")
	case D:
		fmt.Println("D")
	case E:
		fmt.Println("E")
	case F:
		fmt.Println("F")
	case G:
		fmt.Println("G")
	case H:
		fmt.Println("H")
	default:
		return
	}
	
}


