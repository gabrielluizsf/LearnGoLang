package main
//cmd/godoc: deprecate and point to cmd/pkgsite
import (
  "fmt"
  "github.com/gabrielluizsf/LearnGoLang/exercicios/LVL12/ex1/cachorro"
)

func main(){
  dog := 2;
  dogAge := cachorro.Idade(dog);
    fmt.Printf("Meu Cachorro tem %d anos",dogAge);
}
