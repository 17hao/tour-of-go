package main

import (
	"errors"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func f1() error {
	// err := errors.New("f1 err")
	// return err
	return nil
}

func f2() error {
	err := errors.New("f2 err")
	// panic(err)
	return err
	// return nil
}

func f3() error {
	logrus.Info("f3")
	return errors.New("f3 err")
}

func f() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logrus.Errorf("panic, err=%v, stack=%s", panicErr, string(debug.Stack()))
			err = f3()
		}
	}()

	err1 := f1()
	if err1 != nil {
		logrus.Errorf("f1 failed, err=%+v", err1)
		return err1
	}

	err2 := f2()
	if err2 != nil {
		logrus.Errorf("f2 failed, err=%+v", err2)
		return err2
	}

	defer func() {
		if err2 != nil {
			logrus.Errorf("defer err=%v", err2)
			err = f3()
		}

	}()

	return nil
}

func main() {
	// err := f()
	// if err != nil {
	// 	logrus.Errorf("f return error, err=%+v", err)
	// 	return
	// }
}
