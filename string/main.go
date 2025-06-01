package main

import (
	"fmt"
	"strings"
	// "strings"
)

func main(){
	fruits :="apple,banana,mango";
	parts :=strings.Split(fruits, ",")
	fmt.Println(parts)

	str := "one two two three three three"
	count :=strings.Count(str,"two	")
	fmt.Println(count);
}