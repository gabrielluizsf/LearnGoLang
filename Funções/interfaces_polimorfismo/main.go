package main

import (
	"fmt"
)

type myName struct{
  name string
}

type student struct{
  myName
  course     string
}

type professional struct{
  myName
  job        string
}

type human interface{
    hello();
}

func (whoami myName) hello(){
    fmt.Println("Hello",whoami.name);
}

func helloHuman(people human){
  people.hello();
  switch people.(type){
  case student:
    fmt.Println("Estudante de ",people.(student).course);
  case professional:
    fmt.Println("Trabalha como ",people.(professional).job);
  default:
      fmt.Println("[ERROR]: INVALID TYPE");
  }
}

func main(){
  developer := professional{
    myName: myName{
    name: "Dev",
    },
    job: "Developer",
  }
  customer := student{
    myName: myName{name: "Student",},
    course: "CS",
  }

  helloHuman(developer);
  helloHuman(customer);
}
