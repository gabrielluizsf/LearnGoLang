package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)
var waitGroup sync.WaitGroup

func main(){

printNumCPUs();
var counter int64;
total_de_goroutines := 100;
waitGroup.Add(total_de_goroutines);
for indice :=0; indice < total_de_goroutines; indice ++{
	go func(){
		atomic.AddInt64(&counter,1);
		runtime.Gosched();
		fmt.Println("Contador:",atomic.LoadInt64(&counter));
		waitGroup.Done();
	}()
}
waitGroup.Wait();
fmt.Println("Resultado:",counter * int64(total_de_goroutines));

}

func printNumCPUs(){	fmt.Println("[CPUs]:",runtime.NumCPU());}
func printGoroutines(){fmt.Println("[Goroutines]:",runtime.NumGoroutine());}