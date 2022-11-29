package main

import (
	"runtime"
	"fmt"
	"github.com/GabrielLuizSF/how-to-use-packages/hello"
)

func HelloUser(){
user := hello.Whoami();
fmt.Println("Hello", user,"\n Software Desenvolvido em um sistema Linux/amd64\nSistema Atual:",runtime.GOOS,"\nArquitetura:" ,runtime.GOARCH);
}