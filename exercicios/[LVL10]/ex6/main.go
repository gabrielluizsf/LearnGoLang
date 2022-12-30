package main

import "fmt"

func main(){
channel := make(chan int)
  go generate1e2NumbersToChannel(channel)
  for value := range channel{
    fmt.Println("Channel Response: ", value)
  }
}

func generate1e2NumbersToChannel(channel chan int){
  for i := 0; i <= 100; i++{
    channel <- i
  }
  close(channel)
}
