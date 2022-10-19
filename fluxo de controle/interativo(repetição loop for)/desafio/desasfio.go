package desafio

import (
    "fmt"
)

func Desafio(){

    for number := 33; number < 123; number++{
        fmt.Printf("Decimal: %d\tHexadecimal: %#x\tUnicode: %#U\tString: %v\n",number,number,number,string(number));
    }
}
