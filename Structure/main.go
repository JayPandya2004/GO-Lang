package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	Age int
}

type Contact struct{
	Email string
	Mobile int
}

type Address struct{
	Address string
}

type Employee struct{
	Person_Details Person
	Person_Address Address
	Person_Contact Contact

}

func main() {
	person1 := Person{
	FirstName:"Jay",
	LastName:"Pandya",
	Age:21,
	}
	fmt.Println("Details of Jay is:",person1)

	employee1 :=new(Employee)
	employee1.Person_Details.FirstName="Harshil"
	employee1.Person_Details.LastName="Pandit"
	employee1.Person_Details.Age=21

	employee1.Person_Address.Address="Pardi"

	employee1.Person_Contact.Email="harshil1234201@gmail.com"
	employee1.Person_Contact.Mobile=2145785336

	fmt.Println("Details are:",employee1)
}