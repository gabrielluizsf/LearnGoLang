package main

import "fmt"

type veiculo struct{
	portas int
	cor    string
}

type caminhonete struct{
veiculo veiculo
traçãoNasQuatro bool	
}

type sedan struct{
veiculo veiculo
modeloLuxo bool
}

func main(){
	Chevrolet_S10 := caminhonete{
		veiculo: veiculo{
		portas: 4,
		cor:    "prata",
		},
		traçãoNasQuatro: true,
	}
	Audi_A3 := sedan{
    veiculo: veiculo{
		portas: 4,
		cor:    "preto",
	},
	modeloLuxo : true,
}
fmt.Println(Audi_A3);
fmt.Println("Audi A3 Sedan S Line Limited 1.4 TFSI é um modelo luxo:",Audi_A3.modeloLuxo);
fmt.Println(Chevrolet_S10);
fmt.Println("Chevrolet S10 tem tração nas quatro rodas:",Chevrolet_S10.traçãoNasQuatro);
}