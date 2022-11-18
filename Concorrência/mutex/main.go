package main

import (
	"fmt"
	"runtime"
	"sync"
)
var waitGroup sync.WaitGroup
var mutex  sync.Mutex

func main(){

printNumCPUs();
printGoroutines();

counter := 0;
total_de_goroutines := 100;
waitGroup.Add(total_de_goroutines);
for indice :=0; indice < total_de_goroutines; indice ++{
	go func(){
		mutex.Lock();
		value := counter;
		runtime.Gosched();
		value++;
		counter = value;
		mutex.Unlock();
		waitGroup.Done();
	}()
	printGoroutines();
}
waitGroup.Wait();
printGoroutines();
fmt.Println("Resultado:",counter);

}

func printNumCPUs(){	fmt.Println("[CPUs]:",runtime.NumCPU());}
func printGoroutines(){fmt.Println("[Goroutines]:",runtime.NumGoroutine());}