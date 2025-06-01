package main

import "fmt"

func add(a int,b int)(result int){
	result = a+b
	return
}
func main() { //{} braces should start from first line of function declaration if we start it in next line it will show errro
	fmt.Println("Hello")
	var ans = add(2,3)
	fmt.Println(ans)
}