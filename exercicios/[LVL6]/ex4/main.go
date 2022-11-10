package main

import "fmt"

type people struct{
  name      string
  lastName  string
  age       int
}
func(pessoa people) whoami(){
  fmt.Println("Meu nome é:",pessoa.name," ",pessoa.lastName,"\nMinha idade é:",pessoa.age);
}
func main(){
  customer := people{
    name: "Doctor",
    lastName: "Who",
     age: 400,
  } 
  customer.whoami();
}
