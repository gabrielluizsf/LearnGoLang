package main

import "fmt"

func main(){
  channel := make(chan int)
    generateGoroutines(10, channel)
    printChannelResponse(100,channel)
}

func printChannelResponse(number int, channel chan int){
  for i := 0; i < number; i++{
    fmt.Println("Channel Response: ",<- channel)
  }
}

func generateGoroutines(numberOfGoroutines int, channel chan int){
  for i:=0; i < numberOfGoroutines; i ++{
    go func(){
      for value := 0; value < numberOfGoroutines; value ++{
        channel <- value
      }
    }()
  }
}
