package main

import "fmt"

func main() {
	//string is the datatype of key and int is the datatype of value
	studentGrade :=make(map[string]int)
	studentGrade["Jay"] = 100
	studentGrade["Vandan"] = 50
	studentGrade["Ram"] = 80

	fmt.Println("Marks of Jay is:",studentGrade["Jay"])
	// delete(studentGrade,"Vandan")
	// fmt.Println("Marks of Vandan is:",studentGrade["Vandan"])

	value,exist :=studentGrade["Jay"]
	fmt.Println("Value",value)
	fmt.Println("Exist",exist)

	for key,value := range studentGrade{
		fmt.Printf("Marks of %s is %d\n",key,value)
	}

	person :=map[string]int{
		"Alice":100,
		"Bob":200,
	}

	for key,value :=range person{
		fmt.Printf("Value of %s is %d\n",key,value)
	}

}