package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main(){
	fmt.Println("Learning from web Services")
	res,err :=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("Error occured")
		return
	}
	defer res.Body.Close()
	
	data,err :=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error Occured")
		return 
	}
	fmt.Println("Reponse:",string(data))

}