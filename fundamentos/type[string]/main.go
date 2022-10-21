package main

import (
	"fmt"
)

func main(){
	s := "Hello Strings"
	sb := []byte(s)
	fmt.Printf("Valor: %v\nTipo:%T\n\n",s,s,s)
	fmt.Printf("Valor: %v\nTipo:%T\nBinário:%b\n\n",sb,sb,sb)
	
	//UTF8
	for _ ,value := range sb{
		fmt.Printf("Valor: %v\nTipo: %T \nUnicode: %#U \nHexadecimal: %#x\n\n",value,value,value,value)
	}
	for _, v := range s{
		fmt.Printf("Valor: %v\nTipo: %T \nUnicode: %#U \nHexadecimal: %#x\n\n",v,v,v,v)
	}
	//ASCII
	for count := 0; count < len(s); count ++{
		fmt.Printf("Valor: %v\nTipo: %T \nUnicode: %#U \nHexadecimal: %#x\n\n",s[count],s[count],s[count],s[count])
	}

}