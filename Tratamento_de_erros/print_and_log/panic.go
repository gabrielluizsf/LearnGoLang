package print_and_log

import(
  "os"
)

func PANIC(){
  _, err := os.Open("file.txt");
    printPANIC_ERROR(err);
}
func printPANIC_ERROR(err error){
  if err != nil{
    panic(err);
  }
}
