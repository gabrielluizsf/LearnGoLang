package main

import (
	"fmt"
)

func main(){
	fmt.Println("Golang não tem while");
	learnContinue();
}
func clock(){
	for hour := 0; hour <= 12; hour ++{	
		fmt.Println("Hora:",hour,":00");
		for minutes := 0; minutes < 60; minutes ++{
			fmt.Print(hour,":",minutes,"\n");
		}
	}
}
func basicLoop(){
	for x := 0; x < 10; x++{
		fmt.Println(x);
	}
}
func fakeWhile(){
	x := 0;
	for x<10{
		fmt.Println("[X menor que 10]","\tValor de X:", x)
		x++
	}
}

func infiniteLoop(){
	for {
		fmt.Println("Ctrl+C to break infinite for")
	}
}
func endingInfiniteLoop(){
		x := 0
	for{
		if x < 10{
			fmt.Println("[X menor que 10]","\tValor de X:", x);
			x++
		}else{
			fmt.Println("[X é igual a 10]","\tValor de X:",x,"\nBye");
			break
		}
	}
}
func learnContinue(){
	for i := 0; i < 20; i++{
	   if i % 2 !=0{
		continue;
	   }
	   fmt.Println(i);
	}
}
