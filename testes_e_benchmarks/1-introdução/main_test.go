package main

import "testing"

func TestSoma(t *testing.T){
  sumnumbers := soma(1,4,5);
  result     := 10;
  if sumnumbers != result{
    t.Error("Expected: ", result, "\nGot: ",sumnumbers);
  }
}

/*
func TestSomaFail(t *testing.T){
  sumnumbers := soma(1,4,5);
  result     := 42;
  if sumnumbers != result{
    t.Error("Expected: ",result,"Got: ",sumnumbers);
  }
}
*/

func TestMultiplica(test *testing.T){
  multnumbers := multiplica(10,10);
  result      := 100;
  if multnumbers != result{
    test.Error("Expected: ",result,"\nGot: ",multnumbers);
  }
}
