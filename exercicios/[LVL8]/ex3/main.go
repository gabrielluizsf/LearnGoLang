package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	
	err := json.NewEncoder(os.Stdout).Encode(users);
		if err != nil{
			fmt.Println(err);
		}
	james := users[0]
		james.printName().printAge().printSayings();
	miss := users[1]
		miss.printName().printAge().printSayings();
	M := users[2]
		M.printName().printAge().printSayings();

}
func (u user) printName() user {
	fmt.Printf("Name: %s %s\n", u.First,u.Last);
	return u
}

func (u user) printAge() user {
	fmt.Printf("Age: %d\n", u.Age);
	return u
}

func (u user) printSayings() {
	fmt.Println("Sayings:", u.Sayings);
}