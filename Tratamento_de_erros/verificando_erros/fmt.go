package verificando_erros

import (
  "fmt"
)

func FmtPrintlnERROR(message string){
  defaultMessage, err := fmt.Println(message)
    PrintERROR(err);
  fmt.Println(defaultMessage);
}
func PrintERROR(err error){
  if err != nil{
    fmt.Println(err);
  }
}
