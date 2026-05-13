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
Pincode json.Number
}

func main (){

	dir := "./"

	db, err := New(dir,nil)
	if err != nil {
		fmt.Println("Error Error",err)
	}

	employees := []User{
		{"Srijal","25","9847021452","Microsoft",Address{"Bagmati Province","Kathmandu","Nepal", " 44600"}}
		{"Sujan","29","9857034129","Google",Address{"Bagmati Province","Lalitpur","Nepal","44700"}}
		{"Rohan","32","9847034128","Meta",Address{"Lumbini Province","Lamahi","Nepal","22414"}}
		{"Rajesh","33","98050326471","E-sewa",Address{"Gandaki Province","Pokhara","Nepal","33700"}}
		{"Amrit","39","9847032128","CG",Address{"Madhest Province","Janakpur","Nepal","45600"}}
	}

	for _, value := range employees{
		db.Write("users", value.Name, User){
			Name: value.Name,
			Age: value.Age,
			Contact : value.Contact,
			Company: value.Company,
			Address: value.Address,

		}
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error!!",err)
	}

	fmt.Println(records)


	allusers := []User{}
	for _, f:= range records {
employeeFound := User{}
if err := json.Unmarshal([]byte(f), &employeeFound);  err  != nil {
	fmt.Println("Error !!", err)
}

allusers = append(allusers , employeeFound)
	} 

} 