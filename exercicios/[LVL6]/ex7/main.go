package main

import "fmt"

func main(){
  hello := func(name string){
          fmt.Println("Hello",name);
          }
  hello("dev");

}
