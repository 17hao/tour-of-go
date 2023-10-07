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
	// if err := create(&employee{ID: 1, Name: "2023-10-07-1"}); err != nil {
	// 	logrus.Error(err)
	// }

	// ======== read ========
	res, err := queryByIDs([]int64{1})
	if err != nil {
		logrus.Error(err)
	} else {
		fmt.Printf("%+v", res)
	}

	// ======== update ========

	// ======== delete ========
}
