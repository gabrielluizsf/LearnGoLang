package main

import "fmt"

func main(){
  friends := map[string][]string{"Gopher":[]string{"encontrar bugs","estudar bastante"},"Tux":[]string{"nadar"}};

  friends["Mestre dos Magos"] = []string{"sumir"};
  delete(friends,"Tux");
  for name, value := range friends{
    fmt.Printf("\nNome:%v\nHobbie:\n",name);
    for indice, hobbie := range value{
      fmt.Println(indice + 1, ")",hobbie);
    }
  }
}
