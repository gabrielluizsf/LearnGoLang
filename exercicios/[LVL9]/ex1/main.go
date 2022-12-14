package main

import (
	"fmt"
	"sync"
	"runtime"
)

var(
	waitGroup sync.WaitGroup
)

func main(){
	
	addGoroutines(4);
		go helloworld();
		go welcomeUser("root");
		go printNumCPUs();
		go printGoroutines();
	waitGroup.Wait();
}
func addGoroutines(value int){
	waitGroup.Add(value);
}
func helloworld(){
	fmt.Println("Hello World !");
	waitGroup.Done();
}
func printGoroutines(){
	fmt.Println(runtime.NumGoroutine());
	waitGroup.Done();
}
func printNumCPUs(){
	fmt.Println(runtime.NumCPU());
	waitGroup.Done();
}
func welcomeUser(name string){
	fmt.Println("Hello ",name);
	waitGroup.Done();
}