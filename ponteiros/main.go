package main

import (
  "fmt"
somar "pointers/cases"
)

func main(){
  //O ponteiro é o endereço do valor de uma variavel na memória
  somar.Mais_Um(6);
  number := 10;
  numberAddress := &number;
  
  fmt.Println(number);
  *numberAddress ++
  fmt.Println(number);
  fmt.Println(*numberAddress);
  
  name := "";
  
  nameAddress := &name
  
  fmt.Printf("Digite seu nome: ");
  fmt.Scanf("%s",nameAddress);
  fmt.Println("Hello",name);
  
  fmt.Printf("INT:%T  PONTEIRO:%T, STRING:%T PONTEIRO:%T",number,numberAddress,name,nameAddress);
}
