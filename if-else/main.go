package main

import "fmt"



func main() {
	fmt.Print("Enter the Age:")
	var age int
	fmt.Scan(&age)
	if age>18{
		fmt.Println("Eligible to vote")
	}else if age==18{
		fmt.Println("Wait for 1 year")
	}else{
		fmt.Println("Not eligible to vote")
	}

}