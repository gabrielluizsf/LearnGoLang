package main

import "fmt"


func main(){
  number := 10;
  fmt.Println(fatorial_loops(number));
  fmt.Println(fatorial_recursividade(number));

}

func fatorial_recursividade(number int)int{
  if number == 1{
      return number
  }
  return number * fatorial_recursividade(number - 1);
}

func fatorial_loops(number int)int{
  total := 1;
  for number > 1{
    total *= number;
    number--
  }
   return total;
}
