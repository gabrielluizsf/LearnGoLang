package print_and_log

import (
  "log"
  "os"
)

func PaniclnLog(){
   _, err := os.Open("file.txt");
    print_PANICLN_ERROR(err); 
}
func print_PANICLN_ERROR(err error){
  if err != nil{
    log.Panicln(err);
  }
}
