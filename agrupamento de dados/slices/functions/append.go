package functions

import (
    "fmt"
)

func AnexarSlice(slice[]int){
    outraslice := []int{10,11,12,13,14,15};
    slice = append(slice,5,6,7,8,9);
    slice = append(slice, outraslice...);
    fmt.Println("Slice Anexado ao outro:\t",slice);
}
