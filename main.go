package main

import (
	"encoding/json"
	"fmt"
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
		err = db.Write("users",value.Name , User
	{
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

   fmt.Println("Records in users collection:" , record)
}
