package main

import "fmt"
func main(){
  value := valueInteger(20);
  fmt.Println(value);
  fmt.Println(valueInteger_AND_valueString(value));
}

func valueInteger(value int)int{
  return value;
}
func valueInteger_AND_valueString(value int)(string,int){
  return "[Value]:",value
}
