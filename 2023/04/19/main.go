package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func main() {
	//if err := f1(); err != nil {
	//	logrus.Error(err)
	//}
	//
	//fmt.Printf("%+v\n", append(f5(), f6()...))

	tokenized()
}

func f1() (err error) {
	defer f4(err)

	err = f2()
	if err != nil {
		logrus.Error("f2 return error")
	}
	return f3()
}

func f2() error {
	logrus.Info("f2")
	return errors.New("f2 error")
}

func f3() error {
	logrus.Info("f3")
	return errors.New("f3 error")
}

func f4(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

func f5() []int64 {
	return []int64{1, 2}
}

func f6() []int64 {
	return []int64{3, 4}
}
