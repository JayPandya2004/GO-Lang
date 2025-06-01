package main

import (
	"fmt"
)

// Function with proper error handling
func errorHandling(age int) (result string, errorname error) {
	if age > 18 {
		errorname = fmt.Errorf("Error Encountered: age is greater than 18")
		return "", errorname
	}
	result = "Code Executed"
	return result, nil
}

func main() {
	age := 28
	result, _:= errorHandling(age)
	fmt.Println("Result:",result)
}
