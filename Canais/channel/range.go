package channel

import "fmt";

const number = 100;

func loopChannel(n int, s chan <- int){
  for i := 0;  i <= n ; i++{
    s <- i;
  }
  closeINTchannel(s);
}

func RangeChannel(){
  c := make(chan int);

  go loopChannel(number, c);

  printReceiver(c);
}

func printReceiver(receiver <- chan int){
  for value := range receiver{
    fmt.Println("Loading ", value, "%");
    if value == number{
      fmt.Println("[WELCOME]");
    }
  }
}
