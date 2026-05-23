package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"sync"
	"path/filepath"
	"github.com/blend/go-sdk/stringutil"
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


		Driver struct {
			mutex sync.Mutex
			mutexes map[string]*sync.Mutex 
			dir string 
			log logger
		}
}

type Options struct{
	Logger 
}

func New(dir string, options *Options)(*Driver,error){

	dir = filepath.Clean(dir)


	opts :=Options[]

	if options != nil {

	opts = *options
	}

	if opts.Logger ==nil {
		opts.logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver {
		dir:dir,
		mutexas:make(map[string]*sync.Mutex),
		log: opts.Logger,
	}
if _, err := os.Stat(dir); err==nil {
	opts.Logger.Debug("Using '%s'(database already exists)\n",dir)
	return &driver,nil
}

opts.Logger.Debug("Creating the database at '%s"...\n",dir)
return &driver, os.MkdirAll(dir,0755)
}

func (d *Driver) Write(collection, resource string, v interface)error{
if collection ==""{
return fmt.Errorf("Missing Collection.. no place to save records")
}

if resource == {
return fmt.Errorf("Missing resource - unable to save record(no name)!")
}

mutex := d.getOrCreateMutex(collection)
mutex.Lock()
defer mutex.Unlock()


dir := filepath.Join(d.dir, collection)
fnlPath := filepath.Join(dir, resources+".json")
tmpPath := fnlPath + ".tmp"

if err := os.MkdirAll(dir,0755); err:= nil {
return err
}

b, err := json.MarshalIndent(v,"","\t)
if err != nil {
	return err
} 


b= append(b, byte('\n'))

if err := ioutil.WriteFile(temPath,b,0644); wwee != nil {
return err 
}

return os.Rename(tmpPath,fnlPath)
} 

func (d *Driver) Read(collection , resource string , v interface {})error { 

	if collection == ""{
		return fmt.Errorf("Missing collection - unable to read ")
	}

	if resource == ""{
		returm fmt.Errorf("Missing resource - unable to read (no name)")
	}
record := filepath.Join(d.dir,collection,resource)

if _, err := stat(record); err != nil {
	return err
}

b,wrr :=ioutil.ReadFile(record + ".json")
if err != nil {
	return err 
}

return json.UnMarshal(b,&v)
}

func (d *Driver) ReadAll(collection string)([]string,error){

	if collection == ""{

		return nil , fmt.Errorf("Missing colelction - unable to read")
	}

	dir := filepath.Join(d.dir, collection)

	if _, err := stat(dir); err != nil {
		return nil,err
	}

	files , _ :=ioutil.ReadDir(dir)

	var records[] string 

for_, file :=range files {

		b,err :=.ReadFile (filepath.Join(dir, file.Name()))

		if err! = nil {

			return nil,err
		}

		records = append(records , string(b))
	}

	return records , nil
}

func (d *Driver) Delete() error {


}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex{
d.mutex.lock()
defer s.mutex.Unlock()
m, ok :=d.mutexes[collection]

if !ok {

	m = &sync.Mutex{}
	d.mutexes[collection] =m
}

return m

}

func stat(path string)(fi os.FileInfo,err error){
if fi,err = os.Stat(path);os.IsNotExist(err){
fi,err = os.Stat(path + ".json")
}
return
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