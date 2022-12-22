package channel

import (
  "fmt"
  "sync"
)


var WaitGroups sync.WaitGroup

func ConvergeChannel(value int){
  par      := make(chan int);
  impar    := make(chan int);
  converge := make(chan int);

  go sendNumbersToChannelParORImpar(value,par,impar);
  go receiveNumbersParORImpar(par,impar,converge);

  for v := range converge{
    fmt.Println("Response: ",v);
  }

}

func sendNumbersToChannelParORImpar(value int,par chan int, impar chan int){
  for number := 0; number < value; number ++{
    if number % 2 == 0{
      par <- number;
    }else{
      impar <- number;
    }
  }
  close(par);
  close(impar);
}
func receiveNumbersParORImpar(par chan int, impar chan int, converge chan int){
  GoRoutineRangeChannel(par, converge);
  GoRoutineRangeChannel(impar, converge);
  WaitGroups.Wait();
  close(converge);
}

func GoRoutineRangeChannel(valueType chan int,  converge chan int){
  WaitGroups.Add(1);
  go func(){
      for value := range valueType{
        converge <- value
      }
    WaitGroups.Done();
  }()
}
