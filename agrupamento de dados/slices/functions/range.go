package functions

import "fmt"
var (
    frutas = []string{"Uva",}
)

func RangeFrutas(frutas[]string){
    frutas = append(frutas, "Uva", "Mamão");
    frutas = append(frutas[1:2], frutas[4:]...);
    for índice, fruta := range frutas{
    fmt.Println("No Índice",índice,"existe a fruta:",fruta);
    }
}

