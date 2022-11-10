package main

import "fmt"

func main(){
  number1 := counter();
  number2 := counter();

  fmt.Println(number1());
  fmt.Println(number2());


}

func counter() func() int{
  number := 0
  return func()int{
    number ++
    return number
  }
}
