package main

import (
	"fmt"
	"github.com/GabrielLuizSF/LearnGoLang/Canais/canais_direcionais"
  "github.com/GabrielLuizSF/LearnGoLang/Canais/channel"
)

func main(){
	canal := make(chan int);
	/**
	*ERRADO:
	*canal <- 42;
	*fmt.Println(<-canal);
	*/	

	//CERTO:
	go func(){
		canal <- 42;
	}()
	fmt.Println(<- canal);
	channelsBuffer(5, 1);

	canais_direcionais.BiDirectionalChannel();
  channel.RangeChannel();
}

func channelsBuffer(number int, bufferNumber int){
	c := make(chan int, bufferNumber);
	c <- number;
	fmt.Println(<- c);
}
