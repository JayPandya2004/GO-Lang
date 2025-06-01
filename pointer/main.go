package main

import "fmt"

func main(){
	num :=100
	ptr :=&num
	val :=*ptr

	fmt.Println(ptr)
	fmt.Println(val)
}