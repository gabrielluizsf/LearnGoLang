package main

import "fmt";

func main(){
  quit    := make(chan int);
  channel := generate1e2ValuesTOChannel(quit);

  receive(channel, quit);

	fmt.Println("about to exit");
}

func generate1e2ValuesTOChannel(quit chan <- int) <-chan int {
	channel := make(chan int)

  go func(){
	 for i := 0; i <= 100; i++ {
		 channel <- i
	 }
    close(channel);
    quit <- 0;
  }()
	return channel
}

func receive(channel <- chan int, quit chan int){
  for{
  select{
    case value := <- channel:
      fmt.Println(value,"%");
    case <- quit :
      return;
  }
}
}

