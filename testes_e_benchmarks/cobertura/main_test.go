package main

import "testing"

func TestSayHelloForYou(test *testing.T){
  expected := "Hello Gabriel";
  result   := SayHelloForYou("Gabriel");
  if result != expected{
    test.Error("Expected: ",expected, "\nGot: ",result);
  }
}
