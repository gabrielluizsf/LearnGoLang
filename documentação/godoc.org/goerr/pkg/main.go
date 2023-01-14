package goerr

import (
  "fmt"
  "log"
)

func PRINT_DEFAULT_ERROR_MESSAGE(err error){
  if err != nil{
    log.Fatalf("[ERROR]: ",err);
  }
}
