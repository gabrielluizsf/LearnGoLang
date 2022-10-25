package main

import (
    "fmt"
)

func main(){
    numbers := []int{43545167,224234,435463,232564,5,4536356,72432,1234586,349434,2415046}
    for _ , number := range numbers{
        fmt.Println("R$",number);
    }
    fmt.Printf("Tipo do Slice Numbers: %T",numbers);
}
