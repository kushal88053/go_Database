package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
	"go.opentelemetry.io/otel/sdk/resource"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string 
	Age     json.Number
	Contact string
	Company string
	Address Address
}

const Version = "1.0.0"

type (
	Logger  interface {
		Fatal(string , ...interface{})
		Error(String , ...interface{})
		Warn(string, ...interface{}) 
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct 
	{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string 
		log Logger 
	}
)

type Options struct {

   Logger
}


func New(dir string , options Options )(*Driver , error){

	dir = filepath.Clean(dir)

	opts = Options{}

	if opts != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}

	driver := Driver{
		dir:     dir,
		mutexs : make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _ , err := os.Stat(dir); err == nil {

		opts.Logger.Debug("Directory exists:", dir)
		return &driver,  nil ;
		
	}
	opts.Logger.Debug("Creating the database at :%s", dir) ;

	return &driver, os.MkdirAll(dir, 0755) ;
}


func (d *Driver)Write(collection , resource string   ) error{

	if collection == "" {
		return fmt.Errorf("collection name cannot be empty") ;
	}

	if resource == "" {
		return fmt.Errorf("resource name cannot be empty") ;
	}

	mutex := d.getOrCreateMutex(d.mutexes, collection)
	mutex.Lock()

	defer mutex.Unlock() 

	dir := filepath.Join(d.dir, collection)
	finalPath := filPath.Join(dir, resource + ".json")

	tmpPath := finalPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	b , err := json.MarshalIndent(resource, "", "  ")

	if(err != nil){
		return err ;
	}

	b = append(b , byte('\n'))

	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return fmt.Errorf("failed to write to temporary file %s: %w", tmpPath, err)
	}

}

func (d *Driver)Read(collection , resource string , v interface{}) error {

if collection == "" {
		return fmt.Errorf("collection name cannot be empty")
	}
	if resource == "" {
		return fmt.Errorf("resource name cannot be empty")
	}

	mutex := d.getOrCreateMutex(d.mutexes, collection)
	mutex.Lock()
	defer mutex.Unlock()
	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource + ".json")
	info, err := stat(finalPath)
	if err != nil {
		return fmt.Errorf("failed to stat file %s: %w", finalPath, err)
	}
	if info == nil {
		return fmt.Errorf("file %s does not exist", finalPath)
	}
	data, err := os.ReadFile(finalPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", finalPath, err)

	} 

	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("failed to unmarshal data from file %s: %w", finalPath, err)
	}
 
	return nil
}

func (d *Driver)ReadAll()(){

}

func (d *Driver)Delete() error {
}

func stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path+ ".json")

	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return info, nil
}

func (d *Driver)getOrCreateMutex(mutexes map[string]*sync.Mutex, key string) *sync.Mutex {
	if m, ok := mutexes[key]; ok {
		return m
	}
	m := &sync.Mutex{}
	mutexes[key] = m
	return m
}


func main() {
	fmt.Println("Hello, World!")
	dir := "/"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}

	employee := []User{
		{
			Name:    "kushal",
			Age:     "22",
			Contact: "1234567890",
			Company: "Vishnu Tech",
			Address: Address{
				City: "Indore ",

				State:   "MP",
				Country: "India",
				Pincode: "452001",
			},
		},
		{
			Name:    "kushal",
			Age:     "22",
			Contact: "1234567890",
			Company: "TechCorp Tech",
			Address: Address{
				City: "Indore ",

				State:   "MP",
				Country: "India",
				Pincode: "452001",
			},
		},
		{
			Name:    "Shubham",
			Age:     "22",
			Contact: "1234567890",
			Company: "Bajali Tech",
			Address: Address{
				City: "Indore ",

				State:   "MP",
				Country: "India",
				Pincode: "452001",
			},
		},
		{
			Name:    "Ayush",
			Age:     "22",
			Contact: "1234567890",
			Company: "Go explore",
			Address: Address{
				City: "Indore ",

				State:   "MP",
				Country: "India",
				Pincode: "452001",
			},
		},
		{
			Name:    "Vishal",
			Age:     "22",
			Contact: "1234567890",
			Company: "USE expore",
			Address: Address{
				City: "Indore ",
				State:   "MP",
				Country: "India",
				Pincode: "452001",
			},
		},
	}

	for _ , temp = range employee {
		err = db.Write("users",value.Name , User{
		Name : temp.Name,
		Age:  temp.Age,
		Contact: temp.Contact,
		Company: temp.Company,
		Address: temp.Address
	})
}

   record,err :=db.ReadAll("users")

   if err != nil {
	fmt.Println("Error reading records:", err)

   }

   allusers := []User{} 

   for _, record := range records {
	employeeFount := User{}

     if err := json.Unmarshal([]byte(record), &employeeFound); err != nil {	
		fmt.Println("Error unmarshalling record:", err)
	 }

	 allusers = append(allusers, employeeFound)
   }

   fmt.Println("All Users:", allusers) ;

   if db.Delete("users", "kushal") ; err != nil {
	fmt.Println("Error deleting record:", err)
   }



}
