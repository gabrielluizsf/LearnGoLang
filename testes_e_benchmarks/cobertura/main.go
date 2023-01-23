package main

import "fmt"

//Está função  retorna a string que diz Hello para o nome recebido por paramêtro
func SayHelloForYou(name string) string{
    return "Hello "+ name;
}

func main(){
  welcome :=  SayHelloForYou("Gabriel");
  fmt.Println(welcome);
}
