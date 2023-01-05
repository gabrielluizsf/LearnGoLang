package main

import (
  "fmt"
  "time"
)

type SpecialError struct{
  data string
}

func (err SpecialError) Error() string{
  return "[ERROR]: ";
}

func printError(err error){
  if err != nil{
    fmt.Println(time.Now(), err);
  }
}

func main(){
  err := SpecialError{"u1F680"}
  printError(err);
}
