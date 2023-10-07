package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func initLocalDB() {
	var err error
	// https://github.com/go-sql-driver/mysql#timetime-support
	db, err = gorm.Open(mysql.Open("shiqihao:123456@tcp(localhost:3306)/my_db?loc=Local&parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// logrus.Errorf("gorm.Open failed, err=%v, stack=%s", err, string(debug.Stack()))
		panic(err)
	}
}

func getLocalDB() *gorm.DB {
	if db == nil {
		initLocalDB()
	}
	return db
}
