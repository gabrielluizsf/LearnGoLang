package functions

import (
    "fmt"
)

func MultiDimensional(slice[][]int){
    fmt.Println(slice);
    fmt.Println("Valor do Slice 0 na posição 0:\t",slice[0][0]);
    fmt.Println("Valor do Slice 2 na posição 3:\t",slice[2][3]);
}
func InsaneMultiDimensional(multislice[][][][]int){
    fmt.Println(multislice);
}
