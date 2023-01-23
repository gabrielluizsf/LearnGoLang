package errorrr

import "fmt"

//Esta função retorna no console a mensagem customizada e o erro se ocorrer erro
func printCustomMessageAndErrors(message string, err error){
  fmt.Println(message,err);
}
