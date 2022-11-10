package main

import "fmt"

func main(){
    printValue :=  retorna_uma_função();
    result := printValue(30)
    fmt.Println(result);
    fmt.Println(retorna_uma_função()(42));
    
}

func retorna_uma_função() func(float64) float64{
  return func(value float64)float64{
    return  value 
  }
}


