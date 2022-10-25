package main

import "fmt"

func main(){
    numbers := [5]int{7154,2435,45463,4234,7878}
    for _ , number := range numbers{
        fmt.Println("R$",number);
    }
    fmt.Printf("Tipo do Array Numbers: %T",numbers);

}
