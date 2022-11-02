package main

import (
  "fmt"
)

type people struct{
    name string
    age  int
}
type professional struct{
  people
  job     string
}

func main(){

  doctor := professional{
    people: people{name:"Doctor Who",age:450},
    job: "Doctor && Time Lord",
  }
  master := professional{
    people: people{name:"The Master",age:600},
    job: "President && Time Lord",
  }

printStructEmbutido(doctor.people.name,doctor.job,doctor.people.age);
printStructEmbutido(master.people.name,master.job,master.people.age);


}

func printStructEmbutido(name string, job string, age int){
  fmt.Println(name,":\n","Job:",job,"\nAge:",age,"+\n");
}
