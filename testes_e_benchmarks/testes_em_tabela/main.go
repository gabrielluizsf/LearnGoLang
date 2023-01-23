package main

func soma(numbers ...int)int{
  total := 0;
  for _, value := range numbers{
    total += value;
  }
  return total;
}
