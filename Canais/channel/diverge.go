package channel

import(
  "fmt"
  "time"
  "math/rand"
  "sync"
)

func DivergeChannel(number int){
  channel1 := make(chan int);
  channel2 := make(chan int);
  go  sendTOChannel(number, channel1);
  go  rangeMainChannel(channel1,channel2, number);
    for value := range channel2{
      fmt.Println("Response: ",value);
    }
}
func sendTOChannel(number int,channel chan int){
  for i := 0; i < number; i ++{
    channel <- i;
  }
  close(channel);
}
func rangeMainChannel(channel1, channel2 chan int, number int){
  var wg sync.WaitGroup;

  for value := range channel1{
    wg.Add(1);
    go func(number int){
      channel2 <- workLoading(number);
      wg.Done();
    }(value)
  }
  wg.Wait();
  close(channel2);
}
func workLoading(number int)int{
  //ie3 == 1000
  time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)));
    return number;
}

