package main

import (
    "slices-golang/exercicio"
	"slices-golang/functions"
)

func main(){
    functions.RangeFrutas([]string{"Banana","Maçã","Melão"});
    functions.FatiarUmaFatia([]int{1,2,3,4,5});
    functions.Sum([]float64{10.57,1029.,38.10,50.01,0.43});
    exercicio.FakeRange([]int{1,2,3,4,5,6,7,8,9,10});
    functions.AnexarSlice([]int{1,2,3,4});
    functions.ControlSlice([]int{12});
    functions.MultiDimensional([][]int{[]int{1,2,3,4},[]int{5,6,7,8},[]int{9,10,11,12},});
    functions.InsaneMultiDimensional([][][][]int{[][][]int{[][]int{[]int{10,20,30},[]int{40,50,60},},[][]int{[]int{70,80,90},[]int{100}},}});
    functions.ArraySubjacente([]int{10,20,50,40,30},);
}


