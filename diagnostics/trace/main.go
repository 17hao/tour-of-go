package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/trace"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	initDB()

	server := gin.New()
	server.Use(
		gin.Recovery(),
	)
	server.GET("/accounts/:id", getAccountByID)
	//go func() {
	//	err = server.Run(":9000")
	//	if err != nil {
	//		panic(err)
	//	}
	//}()

	err = server.Run(":9000")
	if err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	//<-signalChan
	select {
	case <-signalChan:
		logrus.Info("receive ctrl-c")
		os.Exit(0)
	}
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(mysql.Open("shiqihao:123456@tcp(47.102.157.109:3306)/my_db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		logrus.Error(err)
	}
}

type account struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	City string `json:"city"`
}

func getAccountByID(c *gin.Context) {
	id := c.Param("id")
	var res account
	err := db.Where("id=?", id).First(&res).Error
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
