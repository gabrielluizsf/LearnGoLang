package main

import "fmt"

var (
    value[5]int
    number[7]int
    text[10]string
)

func main(){

    value[0]=1000;
    value[1]=10000000;

    fmt.Println(value[0], value[1], value);
    fmt.Printf("[Array Types]:\nValue: %T\n",value);
    fmt.Printf("Text: %T\n",text);
    fmt.Printf("Number:%T\n",number);

}


