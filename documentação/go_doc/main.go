package main

import (
  "github.com/GabrielLuizSF/go_url/pkg/system/commands"
)

func main(){
  commands.Execute("pwd");
  commands.Execute("go doc fmt.Println");
}

