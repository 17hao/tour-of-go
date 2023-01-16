package main

import "github.com/pkg/errors"

func f1() error {
	return errors.New("f1 error")
}

func f2() error {
	return f1()
}

func f3() error {
	return f2()
}
