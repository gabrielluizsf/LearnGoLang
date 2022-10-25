package functions

import "fmt"

func Profit() {
	//slice
	names := []string{"Fulano"}
	jobs := []string{"QA", "Back-End Developer", "Front-End Developer", "Data Engineer", "DevOps Engineer", "DBA", "SysAdmin"}
	cash := []int{1250, 1370, 1490, 2870, 3279, 4430, 2190}
	//array
	profit := [7]int{150000, 250000, 300000, 40000, 5000, 6000, 7562000}
	cash = append(cash, 5798)
	names = append(names, "Tester")
	fmt.Printf("Name: %v \tJob: %v --> Profit: %v --> Cash: %v\n", names[1], jobs[0], profit[2], cash[7])
	fmt.Printf("Name: %v \tJob: %v --> Profit: %v --> Cash: %v", names[0], jobs[0], profit[4], cash[4])
}
