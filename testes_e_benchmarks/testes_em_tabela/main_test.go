package main

import "testing"

type Teste struct{
  data []int
  answer int
}

func TestSomaEmTabela(test *testing.T){
  tests := []Teste{
      Teste{data: []int{1,4,5}, answer: 10},
      Teste{data: []int{2,5,7}, answer: 14},
      Teste{data: []int{-7,0,7,20},answer: 20},
  }
  for _, value := range tests{
    sum := soma(value.data ...);
    if sum != value.answer{
      test.Error("Expected: ",value.answer,"Got: ",sum);
    }
  }
}
