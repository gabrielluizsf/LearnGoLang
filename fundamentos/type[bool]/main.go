package main

import (
	"fmt"
)

var (
	x bool
)
func main(){
	
	fmt.Println(x)
	
	x  = true
	x  = (1010 == 2022)
	y :=    200 < 1000
	z :=    700 > 2000

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}


