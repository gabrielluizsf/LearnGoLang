package main

import (
	"fmt"
)

type pessoa struct{
	name  string
}

type humanos interface{
	falar()
}

func (people *pessoa) falar(){
	fmt.Println(people.name," diz oi!");
}
func dizerAlgumaCoisa(human humanos){
	human.falar();
}
func main(){
	people := pessoa{"Elliot"};
	people.falar();// É uma shortcut para (&people).falar();
	//imprime people.name diz oi!
	dizerAlgumaCoisa(&people);
	/*
		Não funciona devido não ser um ponteiro para a struct pessoa
	dizerAlgumaCoisa(people);
	*/
}
