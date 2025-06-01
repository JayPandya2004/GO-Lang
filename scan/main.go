package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Print("Enter the name:")
	// fmt.Scan(&name);
	// fmt.Print(name);

	reader :=bufio.NewReader(os.Stdin)
	names, _ := reader.ReadString('\n');
	fmt.Print(names)
}