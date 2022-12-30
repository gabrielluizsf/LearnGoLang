package main

import "fmt";

func main(){
  c := newChannel();
  receive(c);

  printMessage("about to exit");
}

func receive(channel <-chan int){
  for value := range channel{
    printInteger(value);
  }

}

func printInteger(integer int){
  fmt.Println("Response: ",integer);
}

func printMessage(message string){
  fmt.Println(message);
}

func newChannel() <-chan int {
	c := make(chan int)
  go func(){
	  for i := 0; i <= 100; i++ {
		  c <- i
	  }
    close(c);
  }()
	return c;
}

