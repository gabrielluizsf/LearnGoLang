package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarPorIdade []user
func (number ordenarPorIdade) Len()int{
	return len(number);
}

func (number ordenarPorIdade) Less(i,j int)bool{
	return number[i].Age < number[j].Age
}

func (number ordenarPorIdade) Swap(i,j int){number[i],number[j]=number[j],number[i]}

type ordenarPorSobrenome []user
func (sobrenome ordenarPorSobrenome) Len()int{
	return len(sobrenome);
}

func (sobrenome ordenarPorSobrenome) Less(i,j int)bool{
	return sobrenome[i].Last < sobrenome[j].Last
}

func (sobrenome ordenarPorSobrenome) Swap(i,j int){sobrenome[i],sobrenome[j]=sobrenome[j],sobrenome[i]}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(ordenarPorIdade(users));
		fmt.Println("Ordenado por idade:")
			harmonizar(users);
	sort.Sort(ordenarPorSobrenome(users));
		fmt.Println("Ordenado por sobrenome:")
			harmonizar(users);
}
func harmonizar(info []user) {
	for i, value := range info {
		fmt.Println(i + 1, "\tIdade:", value.Age, "\tNome completo:", value.First, value.Last, "\n")
		for _, c := range value.Sayings {
			fmt.Println("\t", c);
		}
		fmt.Printf("\n");

	}
}
