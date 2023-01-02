package main

import (
  "fmt"
  "github.com/GabrielLuizSF/LearnGoLang/Tratamento_de_erros/verificando_erros"
  "github.com/GabrielLuizSF/LearnGoLang/Tratamento_de_erros/print_and_log"
  "github.com/GabrielLuizSF/LearnGoLang/Tratamento_de_erros/recover_example"
  "github.com/GabrielLuizSF/LearnGoLang/Tratamento_de_erros/custom_errors"
)

func main(){
  var err error;
  // valor zero do tipo error == <nil>
    fmt.Println(err);

  verificando_erros.FmtPrintlnERROR("Hello World");
  verificando_erros.ScanERRORS();
  verificando_erros.CreateFile();
  verificando_erros.ReadFile();

  recover_example.Recovered();
  custom_errors.Start();
  print_and_log.Save_Logs();
  print_and_log.PaniclnLog();
}
