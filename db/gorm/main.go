package main

import (
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(string(debug.Stack()))
		}
	}()

	// ======== create ========
	// if err := create(&employee{ID: 1, Name: "2023-10-19-1"}); err != nil {
	// 	logrus.Error(err)
	// }

	// ======== read ========
	res, err := queryByName("a; drop table users")
	if err != nil {
		logrus.Error(err)
	} else {
		fmt.Printf("%+v", res)
	}

	// res, err := queryByTimeRange(time.Now())
	// if err != nil {
	// 	logrus.Error(err)
	// } else {
	// 	for _, e := range res {
	// 		fmt.Printf("%+v\n", e)
	// 	}
	// }

	// ======== update ========

	// ======== delete ========
}
