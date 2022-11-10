package main

import "fmt"

func main(){
  slice := []int{10,20,30,40,50}
  somarNúmeros := sum(slice ...);
  fmt.Println(somarNúmeros);
  somarSlices := sumSlices(slice);
  fmt.Println(somarSlices);
}

func sum(numbers ... int)int{
  result := 0;
   for _ , value := range numbers{
     result += value;
   }
   return result;
}
func sumSlices(numbers []int)int{
  result := 0;
  for _ , value := range numbers{
    result += value;
  }
  return result;

}
