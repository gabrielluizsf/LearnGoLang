package  main

import "fmt"

func main(){
    birthyear := 2001
    for{
        fmt.Println(birthyear);
        birthyear++
        if birthyear == 5000{
            break
        }
    }

}
