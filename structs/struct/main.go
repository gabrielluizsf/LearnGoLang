package main

import "fmt"

type customer struct{
  name    string
  invoice float64
  paid    bool
}

func main(){
  customer1 := customer{
    name:     "customer1",
    invoice:  400.00,
    paid:     true,
  }    
  customer2 := customer{
    name:     "customer2",
    invoice:  1000.00,
    paid:     false,
  }
    isPaid(customer1.name,customer1.paid);
    isPaid(customer2.name,customer2.paid);
}
func isPaid(name string,paid bool){
  var result string;
    if(paid == false){
      result = "Tem que pagar a conta";
    }else{
      result = "Conta Paga";
    }
    fmt.Println(name,":",result);
}
