package main

import "fmt"

func main() {
	number :=[]int{1,2,3,4,5}

	for index,value := range number{
		fmt.Printf("Index:%d have value:%d\n",index,value)
	}
}