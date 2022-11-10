package main

import "fmt"

func main(){
  slice := []int{10,20,30,40,50,60,80}
  defer fmt.Println(sum(slice ...));
  fmt.Println("Result:");
}

func sum(numbers ... int) int{
  result := 0;
  for _ , value := range numbers{
    result += value;
  }
  return result;
}
