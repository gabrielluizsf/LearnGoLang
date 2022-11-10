package main

import "fmt"

func main(){
  number := 10.0;
  numberInput := 0.0;

  print := func(text string, result float64){
    fmt.Println(text,result);
  }
  save_number_input :=  func(){
  fmt.Printf("Digite o número para ele ser multiplicado por 10: ");
  fmt.Scanf("%f"multiplicar(number),&numberInput);
  }

  multiplicar :=  func(ten float64)(string,float64){
  numberOutput := ten * numberInput;
  return "[Resultado]:",numberOutput
  }
  multiplicar(number)
  save_number_input();
  print(multiplicar(number));
  
    
}
