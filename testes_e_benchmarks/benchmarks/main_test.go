package main

import "testing"

type test struct{
  data []int
  answer int
}

func BenchmarkSoma(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
      Soma(10,10);
  } 
}

func BenchmarkMultiplica(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
    Multiplica(50,10)
  }
}
