package exercicio

import "fmt"

func FakeRange(slice[]int){
    for value := 0; value <= len(slice); value++{
        fmt.Println(value);
    }
}
