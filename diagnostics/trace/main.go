package main

import (
	"bufio"
	"encoding/json"
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

// go tool trace trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		logrus.Errorf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Errorf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		logrus.Errorf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	initDB()

	router := gin.New()
	router.Use(
		gin.Recovery(),
	)
	router.GET("/accounts/:id", getAccountByID)
	router.GET("/accounts", getAllAccounts)
	go func() {
		// server.Run() 会阻塞后续代码的执行，所以需要在新的 goroutine 中执行
		err = router.Run(":9000")
		if err != nil {
			panic(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	select {
	case <-signalChan:
		logrus.Info("receive ctrl-c")
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

func getAllAccounts(c *gin.Context) {
	var accounts []account
	err := db.Find(&accounts).Error
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		logrus.Error(err)
		return
	}
	f, err := os.Create("accounts.txt")
	if err != nil {
		logrus.Error(err)
	}
	buf := bufio.NewWriter(f)
	accountsByte, err := json.MarshalIndent(accounts, "", "\t")
	if err != nil {
		logrus.Error(err)
	}
	_, err = buf.Write(accountsByte)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, accounts)
}
