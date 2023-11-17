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

	router := gin.New()
	router.Use(
		gin.Recovery(),
	)
	router.GET("/accounts", getAllAccounts)
	router.GET("/accounts/:id", getAccountByID)
	router.PUT("/accounts/:id", updateAccountByID)
	srv := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(mysql.Open("admin:123456@tcp(localhost:13306)/my_db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
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
