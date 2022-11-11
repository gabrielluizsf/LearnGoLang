package main

import "fmt"

func main(){
  
  count := counter();
  count();
  count();
  number := counter();
  number();
  number();

  
}

  func counter()func(){
    number := 0
    return func(){
        number ++
        fmt.Println(number);
    }
  }

