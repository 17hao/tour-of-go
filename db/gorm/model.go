package main

import "time"

type employee struct {
	ID          int64
	Name        string
	PhoneNumber string
	City        string
	Age         int64
	CreatedAt   time.Time
}
