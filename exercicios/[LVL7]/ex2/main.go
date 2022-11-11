package main

import "fmt";

type people struct{
	name string
	age  int
}

func mudeMe(pessoa *people){
	 (*pessoa).name = "Mudou o nome"
	 (*pessoa).age  =  15
}

func main(){
	test := people{
		name:"test",
		age: 18,
	}
	fmt.Println(test);
	mudeMe(&test);
	fmt.Println(test);
}