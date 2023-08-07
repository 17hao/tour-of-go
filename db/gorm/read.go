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
	db.Model(&employee{}).Where("id in ?", []int64{1, 2}).Find(&es)
	// db.Where("id > ?", "100").Find(&es)
	logrus.Info(db.RowsAffected)
	if db.Error != nil {
		logrus.Error(db.Error)
	}
	return es
}

func queryByID(db *gorm.DB) *employee {
	e := &employee{}
	// if err := db.Where("id=?", "30000").Order("id").Limit(1).Find(&e).Error; err != nil {
	if err := db.Table("employees").Where("id=?", "1").First(e).Error; err != nil {
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

func queryID(db *gorm.DB) []int64 {
	var res []int64
	sql := db.Model(&employee{}).Select("id").Order("id").Find(&res)
	logrus.Info(sql.Statement.SQL.String())
	logrus.Info(db.DryRun)
	err := sql.Error
	if err != nil {
		logrus.Fatalf("%+v\n", err)
	}
	return res
}

func queryFirstID(db *gorm.DB) *employee {
	res := &employee{}

	dryRun := db.Session(&gorm.Session{DryRun: true}).Where("id >= ?", 1).Order("id").First(res)
	logrus.Info(dryRun.Statement.SQL.String())

	err := db.Session(&gorm.Session{DryRun: true}).Where("id >= ?", 1).Order("id").First(res).Error
	if err != nil {
		logrus.Error(err)
		return nil
	}
	// sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return db.Order("id").First(res)
	// })
	// logrus.Info(sql)
	return res
}

func Paginate(db *gorm.DB, limit int, offset int) []employee {
	res := make([]employee, 0)
	err := db.Order("id").Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		logrus.Error(err)
	}
	return res
}

func queryMultiWhere(db *gorm.DB) []employee {
	res := make([]employee, 0)

	db.Where("id >= ?", 1).Where("name like 'tx%'").Find(&res)
	return res
}
