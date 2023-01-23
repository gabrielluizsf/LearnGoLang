package main
/*
package errorrr
Esta função retorna no console a mensagem customizada e o erro se ocorrer erro

func PrintCustomMessageAndErrors(message string, err error){
  fmt.Println(message,err);
}
*/

import (
	"os/user"
	"github.com/theGOURL/errorrr"
)

// this function prints the username of the system,
// returning the username
func main() string {
	whoami, err := user.Current();
  errorrr.PrintCustomMessageAndErrors("ERROR: ",err);
	username := whoami.Username;
	return username;
}
