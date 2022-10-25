package main

import (
    "fmt"
)

func main(){
    slice := []int{23893,82139082309,483480,7428974,281739821,217389723,76437864,37218947,81244,734219847}
    fmt.Println(slice[:3],slice[4:],slice[1:7],slice[2:9]);
    fmt.Println(slice[2:len(slice)-1]);
}
