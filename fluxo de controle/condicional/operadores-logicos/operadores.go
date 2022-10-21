package operadoreslogicos

import (
    "fmt"
)

func LogicOperators(value int){
    if (value % 2 != 0 && value % 3 != 0){
            fmt.Println("É mutiplo de 2 e de 3");
    }
}
