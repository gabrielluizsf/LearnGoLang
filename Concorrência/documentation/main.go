package main

import (
	"fmt"
	start "race-condition-documentation/race_condition"

)



func main(){
	//race condition
	start.RaceCondition();
}


func printDocumentation(){

	readme := `

Simultaneidade - Concorrência

- Compartilhar por comunicação

Programação concorrente é um tópico extenso e há espaço apenas para alguns Destaques específicos do Go aqui.

A programação simultânea em muitos ambientes é dificultada pela sutilezas necessárias para implementar o acesso 
correto às variáveis ​​compartilhadas. Go encoraja uma abordagem diferente na qual os valores compartilhados são 
transmitidos nos canais e, de fato, nunca compartilhado ativamente por threads de execução separadas. Apenas uma goroutine tem acesso ao valor em um determinado momento. As corridas de dados não podem ocorrer, por design. Para encorajar este modo de pensar, nós o reduzimos a um slogan:

    -	Não se comunique compartilhando memória; em vez disso, compartilhe a memória comunicando-se. 

Essa abordagem pode ser levada longe demais. As contagens de referência podem ser melhor realizadas colocando um mutex em torno de uma variável inteira, por exemplo. Mas como um abordagem de alto nível, usar canais para controlar o acesso facilita escrever programas claros e corretos.

Uma maneira de pensar sobre esse modelo é considerar um típico single-threaded programa rodando em uma CPU. Não há necessidade de primitivas de sincronização. Agora execute outra instância; ele também não precisa de sincronização. Agora deixe aqueles dois se comunicam; se a comunicação for o sincronizador, ainda não há necessidade para outra sincronização. Pipelines Unix, por exemplo, se encaixam nesse modelo perfeitamente. Embora a abordagem de concorrência de Go tenha se originado na abordagem de Hoare Comunicação de Processos Sequenciais (CSP), também pode ser visto como uma generalização type-safe de pipes Unix. 
	
[Link] - https://go.dev/doc/effective_go#concurrency
	
	`
	fmt.Println(readme);
}