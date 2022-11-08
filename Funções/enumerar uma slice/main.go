package main

import "fmt"

func main(){
  
  numbers := []int{30,10,20,20,15,5};

  result := sum(numbers ...);

  fmt.Println("O resultado da soma é:",result);



}

func sum(numbers ... int) int{

  result := 0;
  for _ , value := range numbers{
    result += value
  }
  return result


}
