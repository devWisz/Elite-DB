package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type {
		Logger interface {
			Fatal(string,...interface{})
			Error( string , ...interface{})
			 Warn(string, ...interface{}) 
			Info(string, ...interface{})
			Debug(string, ...interface{})
			Trace(string, ...interface{})
		}
}

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

	// so these is the comments added by a human so read this out properly
// for now these code is commented as it works to delete the users so if you want to delete a particular user or all the user uncomment this code 



// Delete a particular user from the database

// 	fmt.Println((allusers))
// if err := db.Delete("user","Srijal") err != nil {
// 	fmt.Println("Error",err)
// }


// Delete all the user from database

// if err := db.Delete("user,""); err != nil {
// fmt.Println("Error",err)
// }
// }