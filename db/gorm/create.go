package main

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func createEmployees(db *gorm.DB, employees []employee) error {
	return db.Create(employees).Error
}

func createEmployee(db *gorm.DB, e employee) error {
	param := []interface{}{e.ID, e.Name, e.City, e.Age}
	return db.Exec("insert into employee(id, name, city, age) values (?, ?, ?, ?)", param...).Error
}

func createOrUpdateOnDup(db *gorm.DB, e employee, name string) error {
	if err := db.Clauses(clause.OnConflict{DoUpdates: clause.Assignments(map[string]interface{}{"name": name})}).Create(&e).Error; err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	return nil
}
