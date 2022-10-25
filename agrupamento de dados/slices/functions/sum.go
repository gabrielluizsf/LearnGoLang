package functions

import "fmt"

func Sum(values[]float64){
    var total float64;
    for _, value := range values{
          total += value;
    }
    fmt.Println("Result: ",total);

}
