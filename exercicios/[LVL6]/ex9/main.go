package main

import "fmt"

func main(){
  helloDev(whoami);
}

func PrintMyName(name string){
  fmt.Printf(name);
}
func whoami(){
  PrintMyName("DEV");
}
func helloDev(f func()){
  fmt.Printf("Hello ");
   f();
}
