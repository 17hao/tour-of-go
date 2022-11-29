package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int64
	name string
	city string
}

func main() {
	db, err := sql.Open("mysql", "shiqihao:123456@tcp(47.102.157.109:3306)/my_db")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	stmtOut, err := db.Prepare("select * from employees where id = ?")
	if err != nil {
		panic(err)
	}
	e := employee{}
	if err := stmtOut.QueryRow(9).Scan(&e.id, &e.name, &e.city); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	fmt.Printf("e=%v\n", e)
}
