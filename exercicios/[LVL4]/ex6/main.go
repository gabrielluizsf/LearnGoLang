package main

import "fmt"

/*
	BRAZIL STATES:
    Acre (AC)
    Alagoas (AL)
    Amapá (AP)
    Amazonas (AM)
    Bahia (BA)
    Ceará (CE)
    Distrito Federal (DF)
    Espírito Santo (ES)
    Goiás (GO)
    Maranhão (MA)
    Mato Grosso (MT)
    Mato Grosso do Sul (MS)
    Minas Gerais (MG)
    Pará (PA)
    Paraíba (PB)
    Paraná (PR)
    Pernambuco (PE)
    Piauí (PI)
    Rio de Janeiro (RJ)
    Rio Grande do Norte (RN)
    Rio Grande do Sul (RS)
    Rondônia (RO)
    Roraima (RR)
    Santa Catarina (SC)
    São Paulo (SP)
    Sergipe (SE)
    Tocantins (TO)

*/

func main(){
	states := make([]string,27,27)
	states = []string{"Acre","Alagoas","Amapá","Amazonas","Bahia","Ceará","Distrito Federal","Espírito Santo","Goiás","Maranhão","Mato Grosso","Mato Grosso do Sul","Minas Gerais","Pará","Paraíba","Paraná","Pernambuco","Piauí","Rio de Janeiro","Rio Grande do Norte","Rio Grande do Sul","Rondônia","Roraima","Santa Catarina","São Paulo","Sergipe","Tocantins"};
	
	fmt.Println("STATES:",len(states),"SLICE CAPACITY",cap(states));
	
	for counter := 0; counter < len(states); counter ++{
		fmt.Println(states[counter]);
	} 
 
}
