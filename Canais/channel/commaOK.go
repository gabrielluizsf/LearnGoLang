package channel

import "fmt";

func CommaOK(number int){
  channel := make(chan int);
  go func(){
    channel <- number;
    close(channel);
  }()

  value, ok := <- channel;

  fmt.Println("Value: ",value,"\nValor Real do Canal? ",ok);
}

 
