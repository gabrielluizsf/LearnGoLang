package main

import "fmt"

func main(){
	somar_pares := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...);
  somar_impares := somenteImpares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...);
	fmt.Println(somar_pares);
  fmt.Println(somar_impares);
}

func soma(numbers ...int) int {
	result := 0
	for _, value := range numbers {
		result += value;
	}
	return result;
}

func somentePares(function func(values ...int) int, numbers ...int) int {
	var slice []int;
	for _, value := range numbers {
		if value%2 == 0 {
			slice = append(slice, value);
		}
	}
	total := function(slice...);
	return total;
}
func somenteImpares(function func(values ...int) int, numbers ...int) int{
  var slice []int;
  for _, value := range numbers {
    if value % 2 != 0{
    slice = append(slice, value);
    }
  }
  result := function(slice ...);
  return result;
}


