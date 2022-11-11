/*
    QUANDO UTILIZAR PONTEIROS:

	1- Quantidades grandes de dados sem ficar copiando para lugares diferentes 
	armazenando em um único lugar na memória.
	2- Teste de Benchmarks.
	3- Mudar valor da posição original.
*/
package cases

import "fmt";

func Mais_Um(number float64){
	counterByValue(number);
	counterByPointer(&number);
}

func counterByValue(number float64){
	number ++;
	fmt.Println(number);

}

func counterByPointer(number *float64){
	*number ++
	fmt.Println(*number);
}