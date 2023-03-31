package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type tb1 struct {
	ID      int64
	Char    string `gorm:"column:char_1"`
	Varchar string `gorm:"column:varchar_1"`
}

func (t *tb1) TableName() string {
	return "table_1"
}

func main() {
	var ss []tb1

	db, err := gorm.Open(mysql.Open("shiqihao:123456@tcp(localhost:3306)/my_db"))
	if err != nil {
		log.Fatal(err)
	}
	db.Find(&ss)

	for _, s := range ss {
		fmt.Printf("%+v\n", s)
	}
}
