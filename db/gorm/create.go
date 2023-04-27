package main

import (
	"gorm.io/gorm"
)

func createEmployees(db *gorm.DB, employees []employee) error {
	return db.Create(employees).Error
}
