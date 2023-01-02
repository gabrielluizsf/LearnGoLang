package print_and_log

import(
  "log"
  "os"
)

func LOG_FATALLN(){
  _, err := os.Open("file.txt");
    printFatalLNERROR(err);
}
func printFatalLNERROR(err error){
  if err != nil{
    log.Fatalln(err);
  }
}
