package main

import "fmt"

func main() {

	anonimo := struct {
		name string
		age  int
	}{
		name: "unknown",
		age:  10000,
	}
	fmt.Println(anonimo);

}
