package main

import "encoding/json"

type Employee struct {
	ID      int
	Name    string
	Address *Address
}

type Address struct {
	Country  string
	Province string
	City     string
}

func main() {
	employee := &Employee{
		ID:   1,
		Name: "test1",
		Address: &Address{
			Country:  "China",
			Province: "ZJ",
			City:     "Hangzhou",
		},
	}
	employeeBytes, _ := json.Marshal(employee)
	println("employee: ", string(employeeBytes))

	employee.Address.City = "Ningbo"
	employeeBytes, _ = json.Marshal(employee)
	println("employee: ", string(employeeBytes))
}
