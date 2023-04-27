package main

import (
	"gorm.io/gorm"
)

func deleteByID(db *gorm.DB, id int64) {
	db.Where("id = ?", id).Delete(&employee{})
}
