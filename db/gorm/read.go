package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func query() employee {
	var e employee
	getLocalDB().First(&e, 3000)
	return e
}

func queryByName(name string) (*employee, error) {
	res := &employee{}
	if err := getLocalDB().Where("name = ?", name).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func queryByTimeRange(startAt time.Time) ([]*employee, error) {
	res := make([]*employee, 0)
	err := getLocalDB().Where("created_at < ?", startAt).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func queryByIDs(ids []int64) ([]employee, error) {
	var es []employee
	res := getLocalDB().Model(&employee{}).Where("id in ?", ids).Find(&es)
	logrus.Info(getLocalDB().RowsAffected)
	if res.Error != nil {
		return nil, res.Error
	}
	return es, nil
}

func queryByID(id int64) (*employee, error) {
	e := &employee{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "table", "employees")
	sql := getLocalDB().WithContext(ctx)
	if err := sql.Where("id=?", id).First(&e).Error; err != nil {
		// if err := getLocalDB().Table("employees").Where("id=?", "1").First(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func queryByAge() []employee {
	var es []employee
	getLocalDB().Where("age=?", "10").Find(&es)
	return es
}

func queryAll() []employee {
	var es []employee
	sql := getLocalDB().Find(&es)
	err := sql.Error
	if err != nil {
		logrus.Fatalf("%+v\n", err)
	}
	for _, e := range es {
		fmt.Printf("%+v\n", e)
	}
	return es
}

func queryByRawSQL() []employee {
	// var e employee
	var es []employee

	getLocalDB().Raw("select id, name, city from employees where id > ?", 1).Scan(&es)
	return es
}

func queryID() []int64 {
	var res []int64
	sql := getLocalDB().Model(&employee{}).Select("id").Order("id").Find(&res)
	logrus.Info(sql.Statement.SQL.String())
	logrus.Info(getLocalDB().DryRun)
	err := sql.Error
	if err != nil {
		logrus.Fatalf("%+v\n", err)
	}
	return res
}

func queryFirstID() *employee {
	res := &employee{}

	dryRun := getLocalDB().Session(&gorm.Session{DryRun: true}).Where("id >= ?", 1).Order("id").First(res)
	logrus.Info(dryRun.Statement.SQL.String())

	err := getLocalDB().Session(&gorm.Session{DryRun: true}).Where("id >= ?", 1).Order("id").First(res).Error
	if err != nil {
		logrus.Error(err)
		return nil
	}
	// sql := getLocalDB().ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return getLocalDB().Order("id").First(res)
	// })
	// logrus.Info(sql)
	return res
}

func Paginate(db *gorm.DB, limit int, offset int) []employee {
	res := make([]employee, 0)
	err := getLocalDB().Order("id").Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		logrus.Error(err)
	}
	return res
}

func queryMultiWhere() []employee {
	res := make([]employee, 0)

	getLocalDB().Where("id >= ?", 1).Where("name like 'tx%'").Find(&res)
	return res
}

func queryMaxAge() []employee {
	res := make([]struct {
		City string
		Age  int64 `gorm:"column:max(age)"`
	}, 2)

	cities := []string{"Hangzhou", "Beijing"}
	getLocalDB().Table("employees").Select("city, max(age)").Where("city in ?", cities).Group("city").Find(&res)
	employees := make([]employee, 0)
	for _, r := range res {
		employees = append(employees, employee{City: r.City, Age: r.Age})
	}
	return employees
}
