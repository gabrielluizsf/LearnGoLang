package main

import (
	"fmt"
	"runtime"
)
var user string;
func main(){

	whoami()
	
	fmt.Println("Hello", user,"\n Software Compilado em um sistema Linux/amd64\nSistema Atual:",runtime.GOOS,"\nArquitetura:" ,runtime.GOARCH);
}
func whoami(){
	fmt.Printf("Qual seu nome: ");
	fmt.Scanf("%s",&user);
}