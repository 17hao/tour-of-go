package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type employee struct {
	ID   int64
	Name string
	City string
}

func main() {
	db, err := gorm.Open(mysql.Open("shiqihao:123456@tcp(47.102.157.109:3306)/my_db"))
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

	e := query(db)
	fmt.Printf("%+v\n", e)
}

func query(db *gorm.DB) employee {
	var e employee
	sql := db.First(&e, 3000)
	fmt.Println(sql.Statement.Table)
	return e
}

func queryAll(db *gorm.DB) []employee {
	var es []employee
	sql := db.Find(&es)
	fmt.Println(sql.Statement.Table)
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

	//db.Raw("select id, name, city from employees where id = ?", 9).Scan(&e)
	//fmt.Printf("e=%+v\n", e)
	//
	//db.Raw("select * from employees").Scan(&es)
	//for _, e := range es {
	//	fmt.Printf("%+v\n", e)
	//}

	//db.First(&e)
	//fmt.Printf("%+v\n", e)

	//db.Last(&e)
	//fmt.Printf("%+v\n", e)
	//
	//db.Find(&es)
	//for _, e := range es {
	//	fmt.Printf("%+v\n", e)
	//}

	//db.Where("city <> ?", "Shanghai").First(&e)
	//fmt.Printf("e = %+v\n", e)

	//db.Where("city <> ?", "Shanghai").Find(&es)
	//for _, i := range es {
	//	fmt.Printf("%+v\n", i)
	//}
}
