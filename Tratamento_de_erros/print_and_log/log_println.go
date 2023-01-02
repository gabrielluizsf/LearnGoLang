package print_and_log

import(
  "log"
  "os"
)

func LOG_PRINTLN(){
  _, err := os.Open("file.txt");
    printLogERROR(err);
}
func printLogERROR(err error){
  if err != nil{
    log.Println("[ERROR]: ",err);
  }
}
