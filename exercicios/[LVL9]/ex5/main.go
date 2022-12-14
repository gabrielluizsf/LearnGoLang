package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	waitGroup sync.WaitGroup
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
		atomic.AddInt64(&counter, 1);
		value := atomic.LoadInt64(&counter);
		runtime.Gosched();
		fmt.Println(value);
		waitGroup.Done();
	}()
	}
}