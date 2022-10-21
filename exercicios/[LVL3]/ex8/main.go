package main

import "fmt"

func main(){
    numberIsValid(10)
}
func numberIsValid(number int){
    name := "Gabriel"
        switch {
        case number < 1000:
            fmt.Println("[VALID]:\t",number," < 1000");
            fallthrough
        case name == "dev":
            fmt.Println("[Creator]:\t",name);
        case number > 1000:
            fmt.Println("[INVALID]",number," > 1000");
        default:
            fmt.Printf("\n[ERROR] Invalid \n")
    }
}
