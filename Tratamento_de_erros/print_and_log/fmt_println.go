package print_and_log

import(
  "fmt"
  "os"
)

func FMT_PRINTLN(){
  _, err := os.Open("file.txt");
    printERROR(err);
}
func printERROR(err error){
  if err != nil{
    fmt.Println("[ERROR]: ",err);
  }
}
