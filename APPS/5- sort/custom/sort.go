package custom

import (
	"fmt"
	"sort"
)


type  Carro struct{
	Nome	  string
	Potencia  int
	Consumo   int
}

type ordenarPorPotencia []Carro

func (number ordenarPorPotencia) Len()int{
	return len(number);
}

func (number ordenarPorPotencia) Less(i,j int)bool{
	return number[i].Potencia < number[j].Potencia
}

func (number ordenarPorPotencia) Swap(i,j int){number[i],number[j]=number[j],number[i]}

type ordenarPorMenorConsumo	[]Carro

func (number ordenarPorMenorConsumo) Len()int{
	return len(number);
}

func (number ordenarPorMenorConsumo) Less(i,j int)bool{
	return number[i].Consumo > number[j].Consumo
}

func (number ordenarPorMenorConsumo) Swap(i,j int){number[i],number[j]=number[j],number[i]}

type ordenarPorMaiorConsumo	[]Carro

func (number ordenarPorMaiorConsumo) Len()int{
	return len(number);
}

func (number ordenarPorMaiorConsumo) Less(i,j int)bool{
	return number[i].Consumo < number[j].Consumo
}

func (number ordenarPorMaiorConsumo) Swap(i,j int){number[i],number[j]=number[j],number[i]}




func MySort(carros []Carro){

fmt.Println("Todos os Carros:\n",carros);

sort.Sort(ordenarPorPotencia(carros));
	fmt.Println("Carro com Menor Potência:\n",carros);
sort.Sort(ordenarPorMenorConsumo(carros));
	fmt.Println("Carro com Menor Consumo:\n",carros);
sort.Sort(ordenarPorMaiorConsumo(carros));
	fmt.Println("Carro com Maior Consumo:\n",carros);
}