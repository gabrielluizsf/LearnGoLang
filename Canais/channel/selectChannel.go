package channel 

import "fmt"

func SelectChannel(number int){
  ch1 := make(chan int);
  ch2 := make(chan int);

  go func(number int){
    for i := 0; i <= number;i ++{
         ch1 <- i;
    }
  }(number/2)

  go func(number int){
    for i := 0;  i <= number; i++{
        ch2 <- i;
    }
  }(number/2)

  for i := 0; i <= number; i ++{
    select{
      case value := <- ch1:
        fmt.Println("Channel 1: ", value);
      case value := <- ch2:
        fmt.Println("Channel 2: ", value);
    }
  }
}

func StartChannel(number int){
  channel := make(chan int);
  quit    := make(chan int);

  go receiveAndQuit(number,channel,quit);
     sendToChannelAndQuit(channel,quit);
}

func receiveAndQuit(number int, channel chan int, quit chan int){
  for i := 0; i <= number; i ++{
    fmt.Println("Channel Response: ", <- channel);
  }
  quit <- 0
}
func sendToChannelAndQuit(channel chan int, quit chan int){
  value := 10;

  for{
    select{
      case channel <- value:
        value++
      case <- quit:
          return

    }
  }

}

func NumbersChannel(number int){
  par   := make(chan int);
  impar := make(chan int);
  quit  := make(chan bool);

  go sendNumbersToChannel(number, par,  impar,  quit);
     solveProblemReceiveNumbersChannel(par,impar,quit);
}

func sendNumbersToChannel(number int,par chan int, impar chan int, quit chan bool){
  for i := 0; i < number; i ++{
    if i % 2 == 0{
      par <- i;
    }else{
      impar <- i;
    }
  }
  close(par);
  close(impar);
  quit <- true;
}

func receiveNumbersChannels(par chan int, impar chan int, quit chan bool){
  select{
  case value := <- par:
    fmt.Println("O número ",value,"é par.");
  case value := <-impar:
    fmt.Println("O número ",value,"é ímpar");
  case <- quit:
    return
  }
}
