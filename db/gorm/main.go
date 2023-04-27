package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
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

	employees := []employee{
		{
			Name: "20230423-1",
		},
		{
			Name: "20230423-2",
		},
	}
	err := createEmployees(db, employees)
	if err != nil {
		logrus.Error(err)
	}
}
