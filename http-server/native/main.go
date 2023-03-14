package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

func getAllAccounts() []account {
	var accounts []account
	err := db.Find(&accounts).Error
	if err != nil {
		logrus.Error(err)
	}
	return accounts
}

func getAccountByID(id int64) (account, error) {
	var res account
	err := db.Where("id=?", id).First(&res).Error
	if err != nil {
		logrus.Error(err)
		return account{}, err
	}
	return res, nil
}

func main() {
	initDB()

	h1 := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		accounts := getAllAccounts()
		accountsBytes, err := json.Marshal(accounts)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(accountsBytes)
	}
	http.HandleFunc("/accounts", h1)

	h2 := func(w http.ResponseWriter, req *http.Request) {
		idStr := strings.TrimPrefix(req.URL.Path, "/accounts/")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		acc, err := getAccountByID(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		accountBytes, err := json.Marshal(acc)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(accountBytes)
	}
	http.HandleFunc("/accounts/", h2)
}
