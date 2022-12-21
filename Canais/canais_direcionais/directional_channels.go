package canais_direcionais

import "fmt"

func BiDirectionalChannel(){
	channel := make(chan int);

	go send(channel, 7);
	receive(channel);
}

func send(s chan <- int, number int){
	s <- number;
}

func receive(r <- chan int){
	fmt.Println("O valor recebido do canal foi: ", <- r);
}
