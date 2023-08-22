package main

import (
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(string(debug.Stack()))
		}
	}()
	db := initDB()

	// fmt.Printf("%+v\n", queryByIDs(db))
	//
	// fmt.Printf("%+v\n", queryByID(db))
	//
	// es := queryByAge(db)
	// fmt.Printf("%+v\n", es)
	//
	// updateTimestamp(db)
	//
	// deleteByID(db, 30000)

	// employees := []employee{
	// 	{
	// 		Name: "20230423-1",
	// 	},
	// 	{
	// 		Name: "20230423-2",
	// 	},
	// }
	// err := createEmployees(db, employees)
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// ids := queryID(db)
	// fmt.Printf("%+v", ids)
	//
	// err := createEmployee(db, employee{ID: 1, Name: "test"})
	// fmt.Printf("%v", err)

	// fmt.Printf("%+v", queryByIDs(db))

	// tx1(db, employee{Name: "tx1"})
	// tx1(db, employee{Name: "tx1"})
	// tx3(db, employee{Name: "tx3"})
	// tx4(db, []*employee{{Name: "tx4-1"}})
	// res := Paginate(db, 1, 2)
	// fmt.Printf("%+v", res)

	// count := int64(0)
	// db.Raw("select * from employees").Count(&count)
	// fmt.Printf("%d", count)
	//
	// res := queryByID(db)
	// fmt.Printf("%+v", res)

	res := queryMaxAge(db)
	fmt.Printf("%+v", res)
}
