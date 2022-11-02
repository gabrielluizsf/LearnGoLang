package main

import "fmt"

type Link struct{
  url string
  websiteName string
}

func main(){
  documentation := Link{
    url: "https://go.dev/doc/effective_go" ,
    websiteName:"Golang Docs",
  }

  printUrl(documentation.url);

}
func printUrl(url string){
   fmt.Println(url)
}
