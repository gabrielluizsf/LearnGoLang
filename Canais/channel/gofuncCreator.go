package channel

import(
  "fmt"
  "time"
  "math/rand"
  "sync"
)

func GoFuncCreator(number_Of_GO_Functions int){
  channel1 := make(chan int);
  channel2 := make(chan int);
  go  sendTOGoChannel(number_Of_GO_Functions, channel1);
  go  rangeMainGoChannel(channel1,channel2, number_Of_GO_Functions);
    for value := range channel2{
      fmt.Println("Response: ",value);
    }
}
func sendTOGoChannel(number int,channel chan int){
  for i := 0; i < number; i ++{
    channel <- i;
  }
  close(channel);
}
func rangeMainGoChannel(channel1, channel2 chan int, number_Of_GO_Functions int){
  var wg sync.WaitGroup;

  for i := 1; i <= number_Of_GO_Functions; i++{
    wg.Add(1)
    go func(){
      for value := range channel1{
        channel2 <- goWorkLoading(value);
      }
      wg.Done();
    }()
  }
    wg.Wait();
    close(channel2);
}
func goWorkLoading(number int)int{
  //ie3 == 1000
  time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)));
    return number;
}

