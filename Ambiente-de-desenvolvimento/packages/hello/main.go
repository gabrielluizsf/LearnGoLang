package hello

import (
	"fmt"
)
var user string;
func Whoami() string{
	fmt.Printf("Qual seu nome: ");
	fmt.Scanf("%s",&user);

	return user;
}