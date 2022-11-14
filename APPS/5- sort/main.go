package main

import (
	"fmt"
	"sort"
)

func main(){

frase   := []string{"Dia","Bom"};
numbers := []int{10,80,90,27,2,5,1,4,3,6,7,8,9,30};
myNumber := 4;

fmt.Println(frase);
fmt.Println(numbers);
SearchMyNumber(numbers,myNumber);

//func Sort(data interface)
sort.Strings(frase);
sort.Ints(numbers);

	fmt.Println(frase);
	fmt.Println(numbers);
	SearchMyNumber(numbers,myNumber);
}
func SearchMyNumber(numbers []int,myNumber int){
//func Sort(data interface)
	search := sort.Search(len(numbers), func(search int) bool { return numbers[search] >= myNumber })
	if search < len(numbers) && numbers[search] == myNumber {
		fmt.Printf("found %d at index %d in %v\n",myNumber, search,numbers)
	} else {
		fmt.Printf("%d not found in %v\n",myNumber,numbers);
	}
}