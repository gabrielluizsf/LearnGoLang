package verificando_erros

import (
  "fmt"
  "os"
  "io"
  "strings"
)
func CreateFile(){
  create, err := os.Create("names.txt");
    PRINT_ERROR_AND_RETURN(err);
  defer create.Close();
  
  readerFile := strings.NewReader("Gabriel Luiz");
  io.Copy(create, readerFile);
}

func PRINT_ERROR_AND_RETURN(err error){
  if err != nil{
    fmt.Println(err);
    return;
  }
}
