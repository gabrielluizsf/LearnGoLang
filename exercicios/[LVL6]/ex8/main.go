package main

import "fmt"

func main(){
  hi := hello();

  hi();

}

func hello() func(){

  return func(){
    fmt.Println("Hello");
  }
}
