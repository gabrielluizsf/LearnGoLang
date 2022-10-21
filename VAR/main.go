package main

import "fmt"

var(
	level = 200
)
func main(){
	level := "LVL "
	print(level)

}

func print(value string){
	fmt.Printf(value)
	fmt.Print(level)
}