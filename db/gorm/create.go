package main

import (
	"gorm.io/gorm"
)

func createEmployees(db *gorm.DB, employees []employee) error {
	return db.Create(employees).Error
}

func createEmployee(db *gorm.DB, e employee) error {
	param := []interface{}{e.ID, e.Name, e.City, e.Age}
	return db.Exec("insert into employee(id, name, city, age) values (?, ?, ?, ?)", param...).Error
}
