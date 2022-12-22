package channel

import "fmt";

func solveProblemReceiveNumbersChannel(par chan int, impar chan int, quit chan bool){
  for {
    select{
    case value := <- par:
      fmt.Println("O número ", value, "é par");
    case value := <- impar:
      fmt.Println("O número", value, "é impar");
    case value, ok := <- quit:
      if !ok{
        fmt.Println("ERROR: ", value);
      }else{
        fmt.Println("Response: ", value, "\nChannel OFF");
      }
      return
    }
  }
}
