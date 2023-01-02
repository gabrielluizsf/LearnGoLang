package custom_errors

import (
  //"errors"
  "fmt"
  "log"
)

type NorgateMathError struct{
  lat string
  long string
  err error
}

func (norgateMatherr NorgateMathError) Error() string{
  return fmt.Sprintf(`A Norgate Math Error Occurred: %v %v %v`,norgateMatherr.lat,norgateMatherr.long,norgateMatherr.err);
}

func Start(){
  _, err := sqrt(-10.79);
    if err != nil{
    log.Println(err);
  }
}
func sqrt(number float64) (float64,error){
  if number < 0{
    norgateMatherr := fmt.Errorf("\nNorgate Math Redux: Square root of negative number: %v",number);
    return 0, NorgateMathError{"50.22983 N", "495478298 W", norgateMatherr};
  }
  return 42, nil;
}

/**
func Start() {
  _,err := sqrt(-10.37);
    if err != nil{
      log.Fatalln(err);
   }
}

func sqrt(number float64) (float64, error){
  if number < 0{
    NorgateMathError := fmt.Errorf("Norgate Math Error: Square root of negative number: %v", number);
    return 0, NorgateMathError;
  }
  return 42, nil;
}
**/


/**
func Start() {
  _,err := sqrt(-10.37);
    if err != nil{
      log.Fatalln(err);
   }
}

func sqrt(number float64) (float64, error){
  if number < 0{
    return 0, fmt.Errorf("Norgate Math Error: Square root of negative number: %v", number);
  }
  return 42, nil;
}
**/

/**
var NORGATE_MATH_ERROR = errors.New("Norgate Math Error: Square root of negative numbers");

func Start(){
  fmt.Printf("%T\n", NORGATE_MATH_ERROR);
  _, err := sqrt(-10);
    if err != nil{
    log.Fatalln(err);
   }

}
func sqrt(number float64) (float64, error){
  if number < 0 {
    return 0, NORGATE_MATH_ERROR;
  }
  return 42, nil;
}
**/

/**
func Start(){
  _, err := sqrt(-10);
    if err != nil{
      log.Fatalln(err);
    }
}
func sqrt(number float64) (float64, error){
  if number < 0{
    return 0, errors.New("Norgate Math: Square root of negative number");
  }
  return 42, nil;
}

**/
