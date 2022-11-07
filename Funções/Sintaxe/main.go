package main

import "fmt"

func main(){
	letras := "abcdefghijklmnopqrstuvwxyz";
	fmt.Println("Letras:",len(letras));
	total,quantidade_de_numeros := sum(10,10,16,28,100);
	fmt.Println("Números Adicionados:",quantidade_de_numeros,"Resultado da soma:", total);
}

func semArgumento(){
	fmt.Println("Oi Bom Dia");
}

func comArgumento(argumento string){
	if argumento == "dia"{
	  semArgumento()
	}else if argumento == "tarde"{
		fmt.Println("Boa tarde");
	}else{
		fmt.Println("Boa noite");
	}
}

func  countLetters(name string){
	fmt.Println("O Nome", name, "tem",len(name),"letras.");
}

func sum(numbers ...int) (int,int){
	result := 0;
	for _ ,value := range numbers{
		result += value;
	}

	return result, len(numbers);

}