package main

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello")
	time.Sleep(2000*time.Millisecond)
	fmt.Println("Hello after pause")
}

func SayHi(){
	fmt.Println("Hi ")
}

func main() {
	go SayHello()
	SayHi()

	time.Sleep(3000*time.Millisecond)
}