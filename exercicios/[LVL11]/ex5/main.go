package main

import "fmt"

const spanish = "spanish"
const portugueseBrazil = "portugueseBrazil";
const helloPrefixInPortugueseBrazil = "Olá ";
const helloPrefixInSpanish = "Hola ";
const helloPrefixInEnglish = "Hello ";

func Hello (name string,language string) string{
	if name == ""{
           name = "Test"
	}
	return setlanguage(language) + name
} 

func setlanguage(language string) (prefix string){

     switch language{	 
        case portugueseBrazil:
		prefix = helloPrefixInPortugueseBrazil
        case spanish:
	    prefix = helloPrefixInSpanish
	default:
	    prefix = helloPrefixInEnglish
	}
	return
      }

func main(){
    var name string;
    fmt.Print("[Your Name]: ");
	  fmt.Scanln(&name);
	  fmt.Println(Hello(name,""));
  }
