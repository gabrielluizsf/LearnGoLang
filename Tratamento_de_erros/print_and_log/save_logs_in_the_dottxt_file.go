package print_and_log

import (
  "os"
  "fmt"
  "log"
)

func Save_Logs(){
  createLogFile , err := os.Create("logs.txt");
    if err != nil{
      fmt.Println("[ERROR]: ",err);
    }
  defer createLogFile.Close();
  
  log.SetOutput(createLogFile);

  open , err := os.Open("file.txt");
    if err != nil{
    log.Println("{ERROR}: ", err);
    }
  defer open.Close();

  fmt.Println("CHECK THE {log.txt} FILE IN THE DIRECTORY")

}
