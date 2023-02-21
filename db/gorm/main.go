package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type employee struct {
	ID   int64
	Name string
	City string
}

func main() {
	employees := query()
	for _, e := range employees {
		fmt.Printf("%+v\n", e)
	}
}

func query() []employee {
	db, err := gorm.Open(mysql.Open("shiqihao:123456@tcp(47.102.157.109:3306)/my_db"))
	if err != nil {
		panic(err)
	}

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
