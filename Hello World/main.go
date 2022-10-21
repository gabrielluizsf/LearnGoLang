package main;

import (
	"fmt"
);


func main(){
	Hello("Hello World");
}
func Hello(value string){
	numberOfBytes, errors := fmt.Println(value);
	fmt.Println(numberOfBytes,errors);
}