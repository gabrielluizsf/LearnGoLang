package main

import "fmt"

func main(){

    for number := 10; number <= 100 ; number ++{
        resto := number%4
        if number <= 99{
        fmt.Printf("[Resto de %v / 4 ]:\t%v\n",number,resto);
        }else{
        fmt.Printf("[Resto de %v/4: ]:     %v\nTHE END\n",number,resto)
        }
     }
}
