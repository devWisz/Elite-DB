package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string 
	Age json.Number
	Contact string 
	Company string 
	Address Address 
}


type Address struct {
City string 
Province string 
Country string 
Pincode Json.Number
}

func main (){

fmt.Println("Welcome to Elite DB")

}