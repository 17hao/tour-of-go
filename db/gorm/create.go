package main

import (
	"fmt"

	"gorm.io/gorm/clause"
)

func mcreate(employees []employee) error {
	return getLocalDB().Create(employees).Error
}

func create(e *employee) error {
	param := []interface{}{e.ID, e.Name, e.City, e.Age}
	return getLocalDB().Exec("insert into employees(id, name, city, age) values (?, ?, ?, ?)", param...).Error
}

func createOrUpdateOnDup(e employee, name string) error {
	if err := getLocalDB().Clauses(clause.OnConflict{DoUpdates: clause.Assignments(map[string]interface{}{"name": name})}).Create(&e).Error; err != nil {
		fmt.Printf("%+v", err)
		return err
	}
	return nil
}
