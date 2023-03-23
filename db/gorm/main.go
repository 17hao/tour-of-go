package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type employee struct {
	ID   int64
	Name string
	City string
	Age  int64
}

func main() {
	db, err := gorm.Open(mysql.Open("shiqihao:123456@tcp(47.102.157.109:3306)/my_db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	//employees := queryByRawSQL(db)
	//for _, e := range employees {
	//	fmt.Printf("%+v\n", e)
	//}

	//employees := queryAll(db)
	//for _, e := range employees {
	//	fmt.Printf("%+v\n", e)
	//}

	//e := query(db)
	//e := queryByID(db)
	//fmt.Printf("%+v\n", e)

	es := queryByAge(db)
	fmt.Printf("%+v\n", es)

	updateTimestamp(db)
}

func query(db *gorm.DB) employee {
	var e employee
	db.First(&e, 3000)
	return e
}

func queryByID(db *gorm.DB) employee {
	var e employee
	db.Where("id=?", "3").Find(&e)
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
		log.Fatalf("%+v\n", err)
	}
	for _, e := range es {
		fmt.Printf("%+v\n", e)
	}
	return es
}

func queryByRawSQL(db *gorm.DB) []employee {
	//var e employee
	var es []employee

	db.Raw("select id, name, city from employees where id > ?", 1).Scan(&es)
	return es
}

func updateTimestamp(db *gorm.DB) {
	err := db.Exec("update time_table set timestamp_1 = ? where id = 1", "2023-03-01 10:00:00").Error
	if err != nil {
		log.Fatal(err)
	}
}
