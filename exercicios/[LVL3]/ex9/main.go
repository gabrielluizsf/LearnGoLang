package main

import "fmt"

func main(){

    mySport("");

}

func mySport(esporteFavorito string){

    switch esporteFavorito{
        case "Football":
            fmt.Println("⚽");
        case "Tênis":
            fmt.Println("🎾");
        case "Football Americano":
            fmt.Println("🏈");
        default:
            fmt.Println("Esporte não Cadastrado");


     }
}

