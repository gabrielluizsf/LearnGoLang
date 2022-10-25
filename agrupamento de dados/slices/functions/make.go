package functions

import (
    "fmt"
)

func ControlSlice(slice[]int){

    slice = make([]int,5,10);
    slice[0] = 1;
    slice[1] = 2;
    slice[2] = 3;
    slice[3] = 4;
    slice[4] = 5;
    slice = append(slice, 6);
    slice = append(slice, 7);
    slice = append(slice, 8);
    slice = append(slice, 9);
    slice = append(slice, 10);
    slice = append(slice, 11);

    fmt.Println("Slice:",slice,"Capacidade:",len(slice),"de",cap(slice));


}
