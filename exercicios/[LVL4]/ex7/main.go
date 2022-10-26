package main

import "fmt"

func main(){
	titles := [][]string{[]string{"nome","sobrenome","hobby"}}
	contents := [][]string{[]string{"i75497543671","Gopher","Golang","localizar erros"},[]string{"i465684554643","Moby","Dock","nadar"},[]string{"i597056660464","Tux","Linux","nadar"}}
	
		for _, content := range contents{
		    fmt.Println("[","ID:",content[0],"]","\n");
			for _, title := range titles{
				fmt.Println("1º",title[0],"\t","2º",title[1],"\t","3º",title[2]);
			}
				if content[1] == "Tux"{
					fmt.Println(content[1],"             ",content[2],"          ",content[3],"\n");
				}else if content[1] == "Moby"{
					fmt.Println(content[1],"           ",content[2],"          ",content[3],"\n");
				}else{
					fmt.Println(content[1],"         ",content[2],"        ",content[3],"\n");
					
				}
		
		}
	}


