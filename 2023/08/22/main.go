package main

import (
	"crypto/sha1"
	"encoding/base64"
	"io"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func hash(data string) string {
	h := sha1.New()
	_, _ = io.WriteString(h, data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
	// return string(h.Sum(nil))
}

type Test struct {
	ID   int64
	Hash string
}

func (Test) TableName() string {
	return "test"
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("shiqihao:123456@tcp(localhost:3306)/my_db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// logrus.Errorf("gorm.Open failed, err=%v, stack=%s", err, string(debug.Stack()))
		panic(err)
	}
	return db
}

func main() {
	db := initDB()
	t := &Test{Hash: hash(`{"type":"object","data":{"url":{"type":"text","data":"https://oa.feishu-boe.cn/flow/api/trigger-webhook/b68e0ff395979897ebf2d5b303378914"},"body":null,"picker":null}}`)}
	db.Create(t)
}
