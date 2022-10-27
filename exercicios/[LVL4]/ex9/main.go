package main

import  "fmt"

func main(){
  friends := map[string][]string{"Gopher":[]string{"Encontrar Bugs","Ler livro do C"},"Tux":[]string{"nadar","observar"}}
  
  friends["Mestre dos Magos"] = []string{"sumir"}

  for key , value := range friends{
    fmt.Printf("\nNome:%v\nHobbie:\n",key);
    for indice, hobbie := range value{
      fmt.Println(indice +1,")",hobbie)
    }
  }

}
