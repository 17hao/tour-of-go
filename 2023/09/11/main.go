package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("shiqihao:123456@tcp(localhost:3306)/my_db?loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		logrus.Errorf("gorm.Open failed, err=%v, stack=%s", err, string(debug.Stack()))
		// panic(err)
	}
	conn, err := db.DB()
	if err != nil {
		logrus.Errorf("db.DB failed, err=%v, stack=%s", err, string(debug.Stack()))
		// panic(err)
	}
	conn.SetMaxOpenConns(100)
}

type employee struct {
	ID        int64
	Name      string
	City      string
	Age       int64
	CreatedAt time.Time
}

func monitor() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				sqlDB, err := db.DB()
				if err != nil {
					logrus.Errorf("db.Db() err=%+v", err)
				}
				fmt.Printf("%s stats=%+v\n", time.Now().String(), sqlDB.Stats())
			}
		}
	}()
	// time.Sleep(10 * time.Second)
	// ticker.Stop()
	// done <- true
}

func main() {
	go func() {
		monitor()
	}()

	count := 190
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		num := i
		go func(int) {
			defer wg.Done()
			e := &employee{}
			// sqlDB, err := db.DB()
			// if err != nil {
			// 	logrus.Errorf("%+v", err)
			// }
			if err := db.Begin().Error; err != nil {
				logrus.Errorf("begin failed, err=%+v", err)
			}
			// defer db.Commit()
			if err := db.Where("id = ?", num).First(e).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Printf("num=%d e=%+v\n", num, e)
				}
				// logrus.Errorf("%s num=%d, err=%+v, stats=%+v", time.Now().String(), num, err, sqlDB.Stats())
			}
			time.Sleep(3 * time.Second)
		}(num)
	}
	wg.Wait()
	time.Sleep(100*time.Second)
}
