package main

import "fmt"

func main(){
  friends := map[string][]string{"Gopher": []string{"Encontrar Bugs"},"Tux":[]string{"nadar","sentar no chão"}};

  for key, value := range friends{
    fmt.Printf("\n\nNome: %v\nHobbie:",key);
    for indice, hobbie := range value{
      fmt.Printf("\n%v) %v",indice + 1,hobbie)
    } 
  }

}

