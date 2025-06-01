package main

import (
	"fmt"
	"time"
)

func main(){
	currentTime :=time.Now()
	fmt.Println(currentTime)

	formatted :=currentTime.Format("02/01/2006")
	fmt.Println(formatted)

	layout_str := "02/01/2006"
	dateStr :="10/05/2100"
	formatted_date, _ :=time.Parse(layout_str,dateStr)
	fmt.Println(formatted_date)

	//add 1 day into current date
	new_date := currentTime.Add(24*time.Hour)
	fmt.Println(new_date)

}