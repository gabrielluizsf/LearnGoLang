package main

import "fmt"

func main(){
  sum  := soma(10,20,30,40);
  mult := multiplica(10,50);


  fmt.Println("O resultado da soma é ",sum);
  fmt.Println("O resultado da multiplicação é ",mult);

}
func soma(numbers ...int)int{
  total := 0;
  for _, value := range numbers{
    total += value;
  }
  return total;
}

func multiplica(numbers ...int)int{
  total := 1;
  for _, value := range numbers{
    total = total * value;
  }
  return total;
}
