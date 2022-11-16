package main

import "fmt"

func main(){

readme := `

Concorrência:
Executar varios processos de forma concorrente com varios processadores ou com apenas um

Multi-Core:
1 processo: -------------->>
wait 1 processo 
2 processo:-------------->>finish
wait 2 processo finish
1 processo ---------------->>finish
wait 1 processo finish
3 processo:----------------->>finish

Single-Core:
1 processo: ---->>finish
wait 1 processo finish
2 processo: -->>finish
wait 2 processo finish
3 processo: --->>finish

Paralelismo:
Executar processos de forma paralela com varios processadores
1 processo: ----------->>finish
2 processo: -->>finish
3 processo: ---------->>finish
4 processo:-->>finish
`
fmt.Println(readme);
}