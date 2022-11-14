package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){

	senha := "teste123";
	senhaerrada := "123teste";

	senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(senha), 10);
	if err != nil {
		fmt.Println(err);
	}

	fmt.Println("Senha Criptografada:",string(senhaCriptografada));

	fmt.Println("Senha Certa:",bcrypt.CompareHashAndPassword(senhaCriptografada, []byte(senha)));
	fmt.Println("Senha Incorreta:",bcrypt.CompareHashAndPassword(senhaCriptografada, []byte(senhaerrada)));

}