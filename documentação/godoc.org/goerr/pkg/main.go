package goerr

import (
  "fmt"
  "log"
)

func PRINT_DEFAULT_ERROR_MESSAGE(err error){
  log.Fatalf("[ERROR]: ",err);
}
