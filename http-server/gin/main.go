package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	initDB()

	server := gin.New()
	server.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/accounts/300"),
		gin.Recovery(),
	)
	server.GET("/accounts", getAllAccounts)
	server.GET("/accounts/:id", getAccountByID)
	server.PUT("/accounts/:id", updateAccountByID)
	err := server.Run(":9000")
	if err != nil {
		panic(err)
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

func getAllAccounts(c *gin.Context) {
	var accounts []account
	err := db.Find(&accounts).Error
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, accounts)
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

func updateAccountByID(c *gin.Context) {
	id := c.Param("id")
	var acc account
	err := c.Bind(&acc)
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		logrus.Error(err)
	}
	err = db.Where("id=?", id).UpdateColumns(acc).Error
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, acc)
}
