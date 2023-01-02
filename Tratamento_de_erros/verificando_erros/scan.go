package verificando_erros

import "fmt"

func ScanERRORS(){
  var answer1, answer2, answer3 string;

  fmt.Print("Name: ");
    scanValueAndVerifyError(answer1);
  fmt.Print("Fav Food: ");
    scanValueAndVerifyError(answer2);
  fmt.Print("Fav Sport: ");
    scanValueAndVerifyError(answer2);

  fmt.Println(answer1,answer2,answer3);
}
func scanValueAndVerifyError(answer string){
  _,err := fmt.Scan(&answer);
   PANIC_ERROR(err); 
}

func PANIC_ERROR(err error){
  if err != nil{    
    panic(err);
  }
}
