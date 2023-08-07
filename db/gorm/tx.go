package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func tx1(db *gorm.DB, e employee) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&e).Error; err != nil {
			return err
		}
		time.Sleep(1 * time.Minute)
		return nil
	})
	if err != nil {
		logrus.Errorf("err=%v", err)
	}
}

func tx2(db *gorm.DB, e employee) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Create(&e).Error; err != nil {
			return err
		}
		time.Sleep(1 * time.Minute)
		return nil
	})
	if err != nil {
		logrus.Errorf("err=%v", err)
	}
}

func tx3(db *gorm.DB, e employee) {
	tx := db.Begin()
	err := tx.Create(&e).Error
	defer func(error) {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}(err)
}

func tx4(db *gorm.DB, e []*employee) {
	tx := db.Begin()
	err := tx.Create(&e).Error
	defer func(error) {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}(err)
}
