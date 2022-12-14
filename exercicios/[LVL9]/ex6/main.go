package main

import (
	"fmt"
	"runtime"
)



func main(){
fmt.Println(myOS(),myARCH());
}
func myOS()string{return runtime.GOOS}
func myARCH()string{return runtime.GOARCH}