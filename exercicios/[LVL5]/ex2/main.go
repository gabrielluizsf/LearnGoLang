package main

import "fmt";

type people struct{
  nome      string
  sobrenome string
  sabores_de_sorvete []string
}


func main(){
customers := make(map[string]people);

customers["da Silva"] = people{
    nome: "João" ,
    sobrenome:"da Silva" ,
    sabores_de_sorvete: []string{"Napolitano","Flocos"},
  }
  customers["Joana"] = people{
    nome: "Maria",
    sobrenome: "Joana",
    sabores_de_sorvete: []string{"Chocolate","Morango"},
  }
   pedidos(customers);


  }

func pedidos(customers map[string]people){
 for _, cliente := range customers{
 fmt.Printf("\n");
 fmt.Println("[Cliente]:",cliente.nome,cliente.sobrenome);
  for indice , pedido := range cliente.sabores_de_sorvete {
    fmt.Println("Pedido:",indice + 1,"\n","Sorvete:",pedido);
  }
}
}
