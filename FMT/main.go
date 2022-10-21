package main

import (
	"fmt"
)

func main(){
	stringInterpretada := "H\ne\nl\nl\no\t World\n"
	stringNaoInterpretada := `H\ne\nl\nl\no\t World`
	unirStrings := fmt.Sprintln(stringInterpretada,stringNaoInterpretada)

	fmt.Println(unirStrings)
	fmt.Println(stringInterpretada)
	fmt.Println(stringNaoInterpretada)

}