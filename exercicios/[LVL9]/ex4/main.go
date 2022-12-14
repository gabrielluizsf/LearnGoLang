package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	waitGroup sync.WaitGroup
	mutex sync.Mutex
	counter   int64
)
func main(){
	generateGoroutines(50);
	waitGroup.Wait();
	fmt.Println("Contador:",counter);
}

func generateGoroutines(number int){
	waitGroup.Add(number);
	for count := 0; count < number; count ++{
	go func(){
		mutex.Lock();
		value := counter;
		runtime.Gosched();
		value++
		counter = value;
		waitGroup.Done();
		mutex.Unlock();
	}()
	}
}