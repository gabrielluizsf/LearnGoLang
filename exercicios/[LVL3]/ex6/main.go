package main

import "fmt"

func isValidNumber(number int){
    condition := number < 1000
    if number != 0 {
        fmt.Println(number,"< 1000:",condition);
    }
}

func main(){
    isValidNumber(10)
}
