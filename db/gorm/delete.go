package main

import (
	"gorm.io/gorm"
)

func deleteByID(db *gorm.DB, id int64) {
	getLocalDB().Where("id = ?", id).Delete(&employee{})
}
