package main

import(
	"fmt"
	"net/url"
)
func main(){
	fmt.Print("URL Learning")
	myURL := "http://myexample2004:8000/path?key1=value1&key2=value2"

	parsedUrl,err :=url.Parse(myURL)
	if err!=nil{
		fmt.Println("Error Encounterd")
		return
	}

	fmt.Println("Scheme:",parsedUrl.Scheme)
	fmt.Println("HostName:",parsedUrl.Host)
	fmt.Println("Path:",parsedUrl.Path)
	fmt.Println("Query:",parsedUrl.RawQuery)

}