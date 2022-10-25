package functions

import "fmt"

func ArraySubjacente(slice[]int){

    fmt.Println("Slice Criado:",slice);

    segundoSlice := append(slice[:2],slice[4:]...);
    fmt.Println("2º Slice:",segundoSlice);
    fmt.Println("Slice que foi criado foi modificado pelo 2º Slice:",slice)
}
