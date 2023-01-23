package main

import (
  "fmt"
)

func ExampleSoma(){
  fmt.Println(Soma(2,4,4));
  //Output: 10
}

func ExampleSomaVários(){
  fmt.Println(Soma(2,1,4));
  fmt.Println(Soma(10,10,5));
  fmt.Println(Soma(40,30));
  //Output:
  // 7
  // 25
  // 70
}
