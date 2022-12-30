package main

import "fmt"

func main(){

  bufferExample(1);
  auto_execute_Example();
}

func bufferExample(bufferNumber int){
  c := make(chan int,bufferNumber);

	c <- 42;

	fmt.Println(<-c);
}

func auto_execute_Example(){
  c := make(chan int);
  go func(){
    c <- 42;
  }();
  fmt.Println(<-c);
}

