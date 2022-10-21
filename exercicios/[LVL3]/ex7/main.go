package main

import "fmt"

func isValidNumber(number int){
    small := number < 1000;
    big   := number > 1000
    if small{
        fmt.Printf("[VALID]: %v < 1000\n",number);
    }else if big{
        fmt.Printf("[VALID]: %v > 1000\n",number);
    }else{
        fmt.Println("[Invalid]");
    }
}
func main(){
    isValidNumber(3);
}
