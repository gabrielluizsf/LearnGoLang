package main

import "fmt"

type people struct{
  name string
}

func (customer people) oibomdia(){
    fmt.Println("Bom dia",customer.name);
}

func main(){
  customer := people{"Customer"};
  customer.oibomdia()

}
