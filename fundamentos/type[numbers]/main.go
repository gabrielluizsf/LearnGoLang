package main

import (
	"fmt"
	"runtime"
)
/*
int(32,64) 
uint(32,64)
float(32,64)
byte == uint8
rune == int32 (UTF8)
*/
func main(){
	
	x := 10
	y := 10.0
	fmt.Printf("X:%T\nY:%T\n",x,y)
	myComputer()
	
}
func printBytes(){
	text   :=  "a z"
	bytesOfText := []byte(text)
	fmt.Printf("Texto: %v \nBytes do Texto: %v",text,bytesOfText)
}
func myComputer(){
	fmt.Println("OS:",runtime.GOOS)
	fmt.Println("Arquitetura:",runtime.GOARCH)
}