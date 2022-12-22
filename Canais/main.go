package main

import (
	"fmt"
	"github.com/GabrielLuizSF/LearnGoLang/Canais/canais_direcionais"
  "github.com/GabrielLuizSF/LearnGoLang/Canais/channel"
  "github.com/GabrielLuizSF/LearnGoLang/Canais/channel/context"
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

  channel.StartChannel(5000);
  channel.RangeChannel();
  channel.SelectChannel(5000);
  channel.NumbersChannel(50);
  channel.CommaOK(55);
  channel.ConvergeChannel(500);
  channel.DivergeChannel(100);
  channel.GoFuncCreator(50);
  context.MainContext();
  channel.WorkChannel(100);
}

func channelsBuffer(number int, bufferNumber int){
	c := make(chan int, bufferNumber);
	c <- number;
	fmt.Println(<- c);
}
