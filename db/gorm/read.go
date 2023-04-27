package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func query(db *gorm.DB) employee {
	var e employee
	db.First(&e, 3000)
	return e
}

func queryByIDs(db *gorm.DB) []employee {
	var es []employee
	db.Where("id > ?", "100").Find(&es)
	logrus.Info(db.RowsAffected)
	if db.Error != nil {
		logrus.Error(db.Error)
	}
	return es
}

func queryByID(db *gorm.DB) *employee {
	e := &employee{}
	// if err := db.Where("id=?", "30000").Order("id").Limit(1).Find(&e).Error; err != nil {
	if err := db.Where("id=?", "30000").First(e).Error; err != nil {
		logrus.Printf("%+v\n", err)
	}
	return e
}

func queryByAge(db *gorm.DB) []employee {
	var es []employee
	db.Where("age=?", "10").Find(&es)
	return es
}

func queryAll(db *gorm.DB) []employee {
	var es []employee
	sql := db.Find(&es)
	err := sql.Error
	if err != nil {
		logrus.Fatalf("%+v\n", err)
	}
	for _, e := range es {
		fmt.Printf("%+v\n", e)
	}
	return es
}

func queryByRawSQL(db *gorm.DB) []employee {
	// var e employee
	var es []employee

	db.Raw("select id, name, city from employees where id > ?", 1).Scan(&es)
	return es
}
