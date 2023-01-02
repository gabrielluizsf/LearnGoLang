package verificando_erros

import(
  "fmt"
  "io/ioutil"
  "os"
)

func ReadFile(){
  open, err := os.Open("names.txt");
   PRINT_ERROR_AND_RETURN(err);
  defer open.Close();

  readALL, err := ioutil.ReadAll(open);
    PRINT_ERROR_AND_RETURN(err);
  fmt.Println(string(readALL));

}
