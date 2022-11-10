package main

import "fmt";

func main(){
  number := 10.0;
  numberInput := 0.0;
  
  func(numbers float64){
  fmt.Printf("Digite o número para ele ser multiplicado por 10: ");
  fmt.Scanf("%f",&numberInput);
  
  fmt.Printf("%1.f",numbers);
  fmt.Printf(" x %1.f",numberInput)
  fmt.Printf("\n[Resultado]: %1.f \n",numbers * numberInput);
  }(number);
}
