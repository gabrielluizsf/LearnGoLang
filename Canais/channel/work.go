package channel

import (
  "fmt"
  "time"
  "math/rand"
)


func WorkChannel(number int){
  channel := convergeWorkChannel(work("Worker 1"),work("Worker 2"), work("Worker 3"));

  for value := 0; value < number; value ++{
    fmt.Println(<-channel);
  }
}

func work(s string) chan string{
  channel := make(chan string);
  go func(s string, c chan string){
    for i := 1; ; i ++{
      channel <- fmt.Sprintf(`%v: %v`,s, i);
      //1e3 == 1000
      time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)));
    }
  }(s,channel)
  return channel
}
func convergeWorkChannel(chan1,chan2,chan3 chan string)chan string{
    newChannel := make(chan string);
    NewChannelReceiveChannel(chan1, newChannel);
    NewChannelReceiveChannel(chan2, newChannel);
  return newChannel;
}
func NewChannelReceiveChannel(channel chan string, newChannel chan string){
  go func(){
    for{
      newChannel <- <- channel
    }
  }()
}
