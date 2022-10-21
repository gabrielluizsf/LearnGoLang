package main

import "fmt"


    func main() {
     for number := 65; number <= 90 ; number ++{
                  fmt.Println(number);
         for counter := 0; counter <=3; counter ++{
                  fmt.Printf("%#U \n",number);
         }

     }
}
