package main

import (
	"fmt"
)

func main() {

	entregas := struct {
		vendedor      string
		pedidos     []string
		clientes   map[int]string
	}{
		vendedor:    "João",
		pedidos: []string{"Torta de Maçã", "Pastel de Queijo"},
		clientes: map[int]string{
			1: "José",
			2: "Maria",
		},
	}
	
	fmt.Println("[Vendedor]:",entregas.vendedor,"{","\n","[Cliente]:",entregas.clientes[1],"\n  [Pedido]:",entregas.pedidos[0],"\n [Cliente]:",entregas.clientes[2],"\n  [Pedido]:",entregas.pedidos[1],"\n}");

}
