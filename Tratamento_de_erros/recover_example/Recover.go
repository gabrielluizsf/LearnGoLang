package recover_example

import "fmt"

func Recovered(){
  sendInt();
  isReturnedNormally();
}

func sendInt(){
  defer func(){
    if rec := recover(); rec != nil{
      fmt.Println("Recovered: ", rec);
    }
  }()
  fmt.Println("Calling .....");
    receiveInt(0);
  isReturnedNormally();
}

func receiveInt(integer int){
  if integer > 3{
    fmt.Println("Panicking");
    panic(fmt.Sprintf("%v", integer));
  }
  defer fmt.Println("Defer in the receiveInt function ", integer);
  fmt.Println("Printing in the receiveInt function ", integer);
  receiveInt(integer + 1);
}

func isReturnedNormally(){
  fmt.Println("{Returned Normally}");
}
