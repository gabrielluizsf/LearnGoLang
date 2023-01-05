package main

import (
	"fmt"
	"log"
  "time"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
    fmt.Println(time.Now());
    valueERROR := fmt.Errorf("\n[Error]: %v", f);
      return 0, sqrtError{"34774","65247",valueERROR};
    }
	return 42, nil;
}
