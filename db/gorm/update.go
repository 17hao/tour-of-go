package main

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func updateTimestamp(db *gorm.DB) {
	err := db.Exec("update time_table set timestamp_1 = ? where id = 10000", "2023-03-01 10:00:00").Error
	if err != nil {
		logrus.Fatal(err)
	}
}

func updateName(db *gorm.DB) {
	err := db.Table("employees").Where("id = ?", 1).Update("name", "new-name").Error
	if err != nil {
		logrus.Fatal(err)
	}
}