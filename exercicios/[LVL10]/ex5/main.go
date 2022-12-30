package main

import "fmt";

func main(){
	
  channel := make(chan int);
  number  := 100;

  go func(){
    channel <- number;
  }()

  value, ok :=  <- channel;
    	fmt.Println("Response: ",value,"\nOK? ",ok,"\n");
  close(channel);

  response, ok := <- channel;
    fmt.Println("Response: ", response,"\nOK?",ok);
}
