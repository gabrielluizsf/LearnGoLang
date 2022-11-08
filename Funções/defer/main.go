package main

import "fmt"

func main(){
  defer fmt.Printf(":)\n");
  defer fmt.Printf(".\n");
  fmt.Printf("Hello World");


}


