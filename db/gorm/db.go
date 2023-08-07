package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("shiqihao:123456@tcp(localhost:3306)/my_db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// logrus.Errorf("gorm.Open failed, err=%v, stack=%s", err, string(debug.Stack()))
		panic(err)
	}
	return db
}
